import numpy.random as rd

def pick_rope(pickShort=True, num_iterations=1000):
    expectation_sum = 0
    
    for iter in range(num_iterations):
        # random cut the rope from 0 to 100
        rope = rd.uniform(0, 100)
        # pick short rope: rope that shorter than 50
        shortRope = rope if rope <= 50 else 100 - rope
        # summary all the short rope
        expectation_sum += shortRope
        print(f"pick the short rope for {iter+1} times", end="\r")
        
    print(f"\nLength expectation of short : {expectation_sum/num_iterations:.4f} ")
        
if __name__ == '__main__':
    pick_rope(num_iterations=100000)