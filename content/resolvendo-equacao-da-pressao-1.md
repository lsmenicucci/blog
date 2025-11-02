---
title: Resolvendo a equacão da pressão (1)
publication_date: 2025-11-01
tags: [fortran, fvm-code]
---

Depois de alguma experimentacão sobre resolver equacoes elípticas, resolvi adotar a bilioteca Petsc para tratar da equacão para a pressão no código de fluido-dinâmica em volumes finitos que estou escrevendo. Previamente, segui a literatura por trás do código EULAG e implementei um precondicionador baseado no método ADI, que decompoe o operador da equacão em diferentes direcões e resolve os problemas unidimensionais alternadamente. Segundo a literatura de volumes finitos, esta técnica não generaliza bem quando consideradas grades genéricas. Como decompor um grade em "fibras" que não se interceptam? Como paralelizar isto para vários hardwares? Não vale a pena me debrucar nestes problemas quando o objetivo principal é utilizar esse cógido para estudar sistemas físicos (e encher a minha tese de douturado de artigos!). Me baseando na quantidade de atualizações ativas do repositório (e no design do site) escolhi a Petsc para resolver a equação da pressão. 


## Instalando e linkando

Tudo aqui é só uma sequência mais didática das [instrucões providas no site da biblioteca](https://petsc.org/release/install/).

Para utilizar a biblioteca em um código Fortran, precisamos informar o compilador onde achar as funcões referenciadas no código e os arquivos com cabeçalhos de definicões que serão inclusos no programa. Respectivamente, precisamos passar as flags `-L<diretorio-1> -I<diretorio-2> -lpetsc`, onde a ultima indica ao compilador que queremos linkar um arquivo `libpetsc.so` no executável. Seja lá onde o petsc foi instalado, no seu diretório base devemos encontrar algo como:
```shell
$ ls .
include  lib  nix-support  share
```
O `nix-support` veio pois instalei o pacote por meio de um `shell.nix`. O `include` tem uma estrutura do tipo:
```shell
$ tree include
include
├── petsc
│   ├── finclude
│   │   ├── petscao.h
│   │   ├── petscbag.h
...
├── petscaodef.mod
├── petscao.h
├── petscao.mod
...
```
Perceba que no primeiro nível, temos cabeçalhos de C `*.h` e módulos de fortran `*.mod` ja precompilados. Um outro diretório `finclude` é contém os cabeçalhos que realmente vamos utilizar no nosso programa em Fortran. 

No diretório `lib` encontramos onde o compilador deve achar a biblioteca dinâmica:
```shell
$ tree lib
lib
├── libpetsc.so -> libpetsc.so.3.23.2
├── libpetsc.so.3.23 -> libpetsc.so.3.23.2
├── libpetsc.so.3.23.2
...
```
O argumento `-l<nome>` induz o _linker_ a procurar um arquivo `lib<nome>.so` em algum dos diretórios passados através do `-L`.

Como era de se esperar de uma biblioteca realmente feita para ser utilizada (cujo oposto é surpreendentemente comum no meio científico), seu pacote gera um arquivo `pkgconfig` o que nos permite gerar as flags a serem passadas para o compilador facilmente:
```shell
$ pkg-config --cflags --libs petsc
-I/nix/store/sqi0ya669hw8317fbn0dr7pmk5hdja39-petsc-3.23.2/include -L/nix/store/sqi0ya669hw8317fbn0dr7pmk5hdja39-petsc-3.23.2/lib -lpetsc
```
Esta resolucão programática pode ser utilizada para carregar as flags automaticamente, como é feito no meu `shell.nix`:
```
{
  pkgs ? import <nixpkgs> { },
  lib ? pkgs.lib,
}:

pkgs.mkShell {
  packages = with pkgs; [
    pkg-config
    hdf5-fortran-mpi
    openmpi
    petsc
  ];

  shellHook = ''
  export FPM_FFLAGS="$FPM_FFLAGS $(pkg-config --cflags --libs hdf5_fortran)"
  export FPM_FFLAGS="$FPM_FFLAGS $(pkg-config --cflags --libs petsc)"
'';
}
```
A variável `FPM_FFLAGS` é simplesmente substituída durente a compilacão, algo do tipo `gfortran $FPM_FFLAGS in.f90 -o out`. Também inclui a blioteca para HDF5, o que não é necessário para o caso atual.

Seguindo o tutorial do site da biblioteca, podemos construir um exemplo mínimo para testar se tudo esta no lugar certo:
```f90
program minimal

#include <petsc/finclude/petscksp.h>
    use petscksp
    implicit none
  
    integer :: ierr

    PetscCallA(PetscInitialize(ierr))
    print *, ierr
    
    PetscCallA(PetscFinalize(ierr))
    print *, ierr

end program
```
Note que o caminho do `#include` é passado relativo a diretório que foi indicado pela flag `-I`.

Compilando com o as flags necessárias e rodando com 1 slot do mpi, esperamos encontrar:
```shell
$ fpm run minimal --runner "mpirun -n 1"
           0
           0

```
O que indica nenhum erro. Vale lembrar que o MPI também foi linkado durente a compilacão.

## Uso geral da biblioteca

Antes de partirmos para um exemplo basico, vamos listar alguns pontos importantes referentes a utilização da biblioteca em um código Fortran comparado ao equivalente em C:
1. Declaracãoes que não são dos tipos `PetscReal` e `PetscInt` podem ser declaradas de duas formas equivalentes:
```f90
type(tXXX) :: b
! ou de forma equivalente 
XXX b
```
O segundo caso utiliza uma macro definida pelo cabeçalho. Por exemplo, um vetor pode ser definido como:
```f90
type(tVec) :: x
! ou de forma equivalente 
PetscVec x
```
2. No caso de variáveis com tipo `PetscReal` e `PetscInt`, a macro resove para um tipo nativo do Fortran, como `real(8)` por exemplo. Nestes casos, só a segunda alternativa é possível.
3. Todas as subrotinas seguem o mesmo nome que a declaração em C mas ao invés de retornarem um inteiro com o código de erro, elas recebem esta variavel como último argumento. Por exemplo, a função `int PetscInitialize()` em Fortran tem a assinatura `PetscFinalize(int)`. Para evitar a escrita de uma condicional checando a variável de erro depois de toda chamada de funcão, a macro `PetscCall` implementa esta checagem:
```shell
$ grep -r "define PetscCall" include/petsc/finclude/*.h
include/petsc/finclude/petscsysbase.h:#define PetscCall(func) call func; CHKERRQ(ierr)
```
A outra macro `CHKERRQ` simplesmente formata o erro e retorna:
```shell
$ grep -r "define CHKERRQ" include/petsc/finclude/*.h
include/petsc/finclude/petscsysbase.h:#define CHKERRQ(ierr) if (ierr .ne. 0) then;call PetscErrorF(ierr,__LINE__,__FILE__);return;endif
```
No programa principal, deve-se utilizar `PetscCallA` que ao invés de retornar quando o erro é não nulo, aborta o MPI.

4. Os módulo que importamos são gerados automaticamente e definem as subrotinas como genéricas. Como as mensagens de erro do gfortran para rotinas genéricas são horriveis, isto dificulta um bocado o processo de descobrir qual parametro da subroutina está sendo passado errado. Um exemplo:
```shell
   37 |     PetscCallA(VecDuplicate(x, b, ierr))
      |                                  1
Error: There is no specific subroutine for the generic ‘vecduplicate’ at (1)
```
Significa que algum destes argumentos esta errado. Se vire pra descrobir qual.

5. Como arrays em C são simplemesmente ponteiros com uma informação de tamanho, rotinas que aceitam arrays são compativeis com escalares passados por referencia. Em Fortran, mesmo que queiremos, por exemplo, inserir um único valor em uma matriz, precisamos passar este valor como um array, da forma `PetscMatSet(..., [valor], ...)`.

## Resolvendo um sistema linear

Vamos resolver o sistema $Ax = b$ como um caso teste:

$$
A = \begin{pmatrix}
1 & 2 & 3 & 4 & 6 \\
6 & 1 & 5 & 3 & 8 \\
2 & 6 & 4 & 9 & 9 \\
1 & 3 & 8 & 3 & 4 \\
5 & 7 & 8 & 2 & 5
\end{pmatrix}
\begin{pmatrix} 1 \\ 2 \\ 3 \\ 4 \\ 5 \end{pmatrix}
=\begin{pmatrix} 60 \\ 75 \\ 107 \\ 63 \\ 76 \end{pmatrix}
$$

O código é auto descritivo, vamos listar alguns comentarios depois:
```f90
program basic 

#include <petsc/finclude/petscksp.h>
    use petscksp
    implicit none

    integer(4), parameter :: n = 5

    type(tVec) :: x, b
    type(tMat) :: A
    type(tKSP) :: ksp
    integer :: i, its
    integer :: ierr

    real(8), dimension(n, n) :: known_A
    real(8), dimension(n) :: known_x, known_b

    ! State the problem: Ax = b
    known_A = reshape([ &
        1, 6, 2, 1, 5, &
        2, 1, 6, 3, 7, &
        3, 5, 4, 8, 8, &
        4, 3, 9, 3, 2, &
        6, 8, 9, 4, 5  &
    ], [5,5])
    known_x = [1, 2, 3, 4, 5]
    known_b = MATMUL(known_A, known_x) ! [60, 75, 107, 63, 76]

    ! Initialize MPI and petsc
    PetscCallA(PetscInitialize(ierr))

    ! Create the x vector and copy its settings to b
    PetscCallA(VecCreate(PETSC_COMM_SELF, x, ierr))
    PetscCallA(VecSetSizes(x, PETSC_DECIDE, n, ierr))
    PetscCallA(VecSetType(x, VECSTANDARD, ierr))

    PetscCallA(VecDuplicate(x, b, ierr))

    ! Set b (the rhs)
    do i = 1, n
        PetscCallA(VecSetValues(b, 1_4, [i - 1], [known_b(i)], INSERT_VALUES, ierr))
    end do

    ! Create the A matrix
    PetscCallA(MatCreate(PETSC_COMM_SELF, A, ierr))
    PetscCallA(MatSetSizes(A, PETSC_DECIDE, PETSC_DECIDE, n, n, ierr))
    PetscCallA(MatSetType(A, MATAIJ, ierr))

    ! Build it
    block
        integer(4) :: cols(n), rows(1)
        cols = [0, 1, 2, 3, 4]

        do i = 1, n
            rows = [i - 1]
            PetscCallA(MatSetValues(A, 1_4, rows, n, cols, known_A(i, :), INSERT_VALUES, ierr))
        end do
    end block

    PetscCallA(MatAssemblyBegin(A, MAT_FINAL_ASSEMBLY, ierr))
    PetscCallA(MatAssemblyEnd(A, MAT_FINAL_ASSEMBLY, ierr))

    ! Create the solver object
    PetscCallA(KSPCreate(PETSC_COMM_SELF, ksp, ierr));
    PetscCallA(KSPSetOperators(ksp, A, A, ierr));

    ! Set the initial value and solve it
    PetscCallA(VecSet(x, -1.0d0, ierr))
    PetscCallA(KSPSolve(ksp, b, x, ierr));

    PetscCallA(KSPGetIterationNumber(ksp, its, ierr));
    print *, "Iterations count: ", its
    print *, "X:"
    PetscCallA(VecView(x, PETSC_VIEWER_STDOUT_WORLD, ierr))

    ! Finalize MPI and petsc
    PetscCallA(PetscFinalize(ierr))

end program
```

Basicamente, montamos um sistema em que sabemos que a solução é `known_x = [1.0, 2.0, 3.0, 4.0, 5.0]`. Com isto construímos o lado direito da equação e depois resolvemos o sistema utilizando um solucionador iterativo fingindo que não sabemos quem é `x`. É instrutivo procurar a assinatura de cada rotina utilizada e comparar com a implementação acima. Antes de verificar a saida, fazemos alguns comentários
1. Selecionamos o _kind_ das variaveis inteiras e reais como `integer(4)` e `real(8)` respectivamente. Esta é a configuração padrão do Petsc, poderíamos ter definido todas as variáveis utilziando `PetscInt` e `PetscReal`. Como o Petsc pode ser configurado utilizando um arquivo externo, depois da compilação, esta última alternativa seria mais portátil do que o código apresentado.
2. Tal como as variávels, utilizamos literais com o _kind_ modificado, como `1_4`. Se erramos o _kind_ a compilação falhará com o erro chato de que a assinatura para o nome genérico não foi encontrada.
3. Alguns argumentos que recebem o tamanho local de objetos receberam o argumento `PETSC_DECIDE`, isto é apenas uma conveniência para subdividir os objetos entre os processos do MPI igualmente.
4. Por padrão, o tipo `tKSP` utiliza o GMRES restartado precondicionado com uma decomposição LU incompleta. 
4. O `block ... end block` é um construtor do Fortran 2008 que apenas nos deixa declarar algums variaveis locais ao bloco antes de executar o código dentro do block. 

Executando, vemos que o código inverte o vetor $X$ corretamente:
```shell
$ fpm run basic --runner "mpirun -n 1"
At line 56 of file app/basic.f90
Fortran runtime warning: An array temporary was created for argument 'f' of procedure 'matsetvalues'
 Iterations count:            1
 X:
Vec Object: 1 MPI process
  type: seq
1.
2.
3.
4.
5.
```
O aviso antes da execução é inofensivo e se deve ao fato que passamos um array que não é contíguo. Em Fortran os arrays são armazenados por coluna. O solucionador convergiu em uma única iteração, provavelmente pois o problema era simples o suficiente para set resolvido pelo precondicionador.
