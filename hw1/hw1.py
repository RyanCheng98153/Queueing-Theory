import random

def play_game(strategy, num_iterations=100000):
    wins = 0
    for iter in range(num_iterations):
        # Place the car behind a random door
        doors = ['goat', 'goat', 'car']
        remaining_doors = [i for i in range(3)]
        random.shuffle(doors)

        # Player chooses a door
        player_choice = random.randint(0, 2)
        remaining_doors.remove(player_choice)
        # Host reveals a goat behind one of the remaining doors
        door_to_reveal = random.choice([i for i in remaining_doors if i != player_choice and doors[i] == 'goat'])
        
        # Player's strategy (change or not change)
        if strategy == 'change':
            remaining_doors.remove(door_to_reveal)
            player_choice = remaining_doors[0]
            
        # Check if player won
        if doors[player_choice] == 'car':
            wins += 1
        
        print(f"[{strategy}]\t win: {wins}, lose: {iter-wins}, wins rate: {wins/(iter+1):.4f}", end='\r')
    print()
    return wins / num_iterations

# Simulate the game with the "not change" strategy
win_rate_not_change = play_game('not change')

if(input("Take a breath ~ Press any key to continue... ") != ''):
    pass
# Simulate the game with the "change" strategy
win_rate_change = play_game('change')

print()
print(f"Probability of winning with strategy 'not change': {win_rate_not_change:.3f}")
print(f"Probability of winning with strategy 'change':\t   {win_rate_change:.3f}\n")