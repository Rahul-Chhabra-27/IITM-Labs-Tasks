### Task

1. You have been given a Person struct.
2. Structure of person
```
Person {
    firstName string
    lastName string
    age int
}
```
3. Your Task is to create a function "changeTheIdentity" that takes two Persons person1 and person2 ( both of them should be pass by reference) as a parameter and swaps the data of both the persons

    Before swapping :-
    person1 = {firstName : 'john',lastName : 'doe;, age :'23'}
    person2 = {firstName : 'james',lastName : 'carry;, age :'30'}
    
    After swapping :-
    person1 = {firstName : 'james',lastName : 'carry;, age :'30'}
    person2 = {firstName : 'john',lastName : 'doe;, age :'23'}

4. Print the persons before calling changeTheIdentity and after the calling changeTheIdentity. 