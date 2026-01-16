## Sunny Sunday

### Instructions

Let's eliminate Sundays by taking them out of the calendar, making a week only six days long, from `"Monday"` to `"Saturday"`.

Create a function called `sunnySunday` that takes a `Date` as an argument and returns the weekday as a `string`.

> 01/01/0001 is a Monday.




Alright — I’ll strip the jargon out and explain it step-by-step like we’re walking through how the program “thinks” in plain terms.
What the program does

This is a Sudoku solver written in Go.
It takes 9 strings (each 9 characters long) as input, where:

    '.' means an empty cell

    '1'–'9' are given numbers

It:

    Checks if the puzzle is valid.

    Makes sure there’s only one possible solution.

    Solves it using backtracking.

    Prints the finished Sudoku.

Breaking it into small, low-level steps
1. isValid – check if a number can go in a spot

    Look through the row → if the number is already there → not allowed.

    Look through the column → if the number is already there → not allowed.

    Find the top-left corner of the 3×3 box the spot belongs to.

    Loop through that 3×3 area → if the number is already there → not allowed.

    If it passes all checks → allowed.

2. solve – the brain

This is recursive backtracking:

    Go cell by cell.

    When you find an empty cell:

        Try numbers '1' to '9'.

        Use isValid to see if the number can fit.

        If it can, put it there and call solve again.

        If the puzzle gets stuck later, remove the number and try the next one.

    If all empty cells are filled → done.

3. countSolutions – check how many solutions exist

    Works like solve, but instead of stopping after 1 solution, it keeps counting.

    If the count reaches limit (here, 2), it stops early — no need to count more.

    Used to make sure there’s exactly 1 solution.

4. copyBoard – make a fresh copy

    Makes a new 9×9 array and copies all numbers from the old board.

    Needed because countSolutions changes the board during testing.

5. printBoard – display the puzzle

    Loops through rows and columns.

    Prints each number (or .) with spaces in between.

6. main – putting it all together

    Check arguments

        Must have exactly 10 arguments (program name + 9 rows).

    Load the board

        Creates a 9×9 array filled with dots.

        Makes sure each row is 9 characters long.

        Only allows digits '1'–'9' or '.'.

    Initial conflict check

        Places each given number only if isValid says it’s okay.

    Check for unique solution

        Uses countSolutions with limit 2.

        If it’s not exactly 1 solution → print "Error".

    Solve the puzzle

        If solve fails → "Error".

    Print result

        Calls printBoard.

In very plain English

Imagine Sudoku as a game board with 81 squares.
This program:

    Reads the puzzle from the command line.

    Makes sure it’s possible to solve (and only one way to solve it).

    Tries numbers in empty spots, one by one.

    Backs up if it makes a wrong guess.

    Stops when the whole thing is filled in correctly.

    Shows you the final filled board.
