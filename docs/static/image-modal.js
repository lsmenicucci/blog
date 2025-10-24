[...document.querySelectorAll('.post .content img')].forEach(img => {
    img.style.cursor = 'pointer';

    img.addEventListener('click', () => {
        const modal = document.getElementById('imgModal');
        modal.querySelector('img').src = img.src;
        modal.classList.add('show');
        document.body.style.overflow = 'hidden'; // lock scroll
    });
});

const imgModalEl = document.querySelector('#imgModal')

if (imgModalEl) {
    imgModalEl.addEventListener('click', e => {
        if (e.target.id === 'imgModal') {
            e.target.classList.remove('show');
        }
        document.body.style.overflow = ''; // unlock scroll
    });
    console.log("Mounted image modal")
}
