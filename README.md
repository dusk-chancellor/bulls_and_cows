# Bulls and Cows Game
![Bulls & Cows](https://repository-images.githubusercontent.com/558048734/9675e878-71d3-4328-9108-5658ed0e69de)
## The game rules:
Random 4-digit number is generated. This number contains different digits from 0 to 9. The main goal is to guess the number. If input number contains a digit from generated number, but not in the right place - it is a cow.  
Else if input number contains a digit from generated number, and is in the right place - it is a bull. The game takes maximum 10 turns, but it may end earlier if guess is right (4 bulls).

e.g.  
generated number is *8413*  
user input: *1402* - *1* bull & *1* cow (as digit *1* is represented in generated number as well as in the user input but in incorrect order, and digit *4* is both represented and in the right place)
### Plan:
- [ ] **random 4-digit number generation**
  - [ ] each digit from 0 to 9
  - [ ] must not contain same digits
- [ ] **user input**
  - [ ] check for valid number (only 4-digit non-repeating number)
  - [ ] check for bulls and cows
- [ ] **multiple user inputs**
  - [ ] the step **user input** is done for 10 times until user guesses the number
  - [ ] if user guesses right -> "YOU WON. The number was ()()()()", or if user fails to guess -> () bulls & () cows
  - [ ] optional: user fails to guess in 10 attempts -> "YOU LOST. The number was ()()()()"
### Further development:
* add interface/make interactive (either framework use or html/css)
* come up with new ideas to improve the game logic


> P.S this is just for practicing skills and fun !
