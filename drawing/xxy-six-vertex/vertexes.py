import matplotlib.pyplot as plt
import matplotlib as mpl
import numpy as np

def draw_h2o_cell(ax, conf, x, y, size, o_neibs = False):
    xs_neib, ys_neib = [], []
    xs_h, ys_h = [], []    
    
    h_idx = 0
    for offset in [0, size]:
        xs_neib.append(x + size/2)
        ys_neib.append(y + offset)
        
        xs_h.append(x + size/2)
        ys_h.append(y + offset/2 + (1 + conf[h_idx])*size/6)        
        h_idx += 1

    for offset in [0, size]:
        xs_neib.append(x + offset)
        ys_neib.append(y + size/2)

        xs_h.append(x + offset/2 + (1 + conf[h_idx])*size/6)        
        ys_h.append(y + size/2)
        h_idx += 1
   
    # Draw o-o lines
    o_line_points = list(zip(xs_neib, ys_neib))
     
    for (axis,(p1, p2)) in enumerate([o_line_points[:2], o_line_points[2:]]):
        line_xs, line_ys = zip(p1, p2)
        
        line_xs = np.array(line_xs)
        line_ys = np.array(line_ys)    
    
        ax.plot(line_xs, line_ys, 'k-', alpha = 0.3) 
        
        if(axis == 1):
            line_xs = x + 0.45*size + 0.1*(line_xs - x)
            ax.plot(line_xs, line_ys + 0.25*size, 'k-', alpha = 0.3) 
            ax.plot(line_xs, line_ys - 0.25*size, 'k-', alpha = 0.3) 
        else:
            line_ys = y + 0.45*size + 0.1*(line_ys - y)
            ax.plot(line_xs + 0.25*size, line_ys, 'k-', alpha = 0.3) 
            ax.plot(line_xs - 0.25*size, line_ys, 'k-', alpha = 0.3) 
    
    ax.plot(xs_h, ys_h, 'ko', mec = '1.0')    

    if (o_neibs == True):
        ax.plot(xs_neib, ys_neib, 'bo', mec = '1.0')
      
    # Draw center oxigen
    ax.plot(x + size/2, y + size/2, 'bo', mec = '1.0')
    

fig = plt.figure()
ax = fig.add_subplot(111)
ax.set_aspect('equal')

draw_h2o_cell(ax, [1, 1, 0, 0], 0, 0, 1, True)

plt.show()        
