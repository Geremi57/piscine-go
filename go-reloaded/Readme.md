## **Step1 - `hex`**

the word b4 replace the word before hex with the decima version of the number to base 10
eg
**1E** -> **30**

## **Step2 - `bin`**

the word bin should replace the word b4 with the decimal version of the word
eg
**10 (bin) years** -> **2 years**

## **Step3 - `up`**

the word up should replace the word b4 it with the uppercase version of the word
eg
**set go(up)! -> **set GO!\*\*

## **Step4 - `low`**

the word low should replace the word b4 it with the lowercase version of the word
eg
**stop SHOUTING(low) -> **stop shouting\*\*

## **Step5 - `cap`**

the word cap should replace the word b4 it with the capitalized version of it
eg
**Brooklyn bridge(cap) -> **Brooklyn Bridge\*\*

## **Step6 - `low,n up,n cap,n**

it turns the previous specified number of words in lowercase uppercase or capitalized
eg
**This is so exciting(up,2) -> This is SO EXCITING**

## **Step7 - `Punctuation rules`**

**/./,/!/?/:/;/**
they shoukd be close to the previous word and with space apart from the next one
eg
**I was sitting over there ,and then BAMM !! -> "I was sitting over there,and then BAMM!!"**

## **Step8 - `gropus of punctuation`**

**.../!?**
in this case the program should be format the text as in the
eg
**I was thinking ... You were right -> I was thinking... You were right**

## **Step9 - `apostrophes`**

**'**
they should always be placed to the end of the words with no spaces between them and a word
eg
**I am exactly how they describe me: ' awesome ' -> I am exactly how they describe me: 'awesome'**

## **Step10 - `a or an`**

**a/an**
instance of a should be turned into an if the next ord begins with a vowel and vice versa if the word begins with a constonant
eg
**There it was, A amzing rock! -> There it was. An amazing rock!**
