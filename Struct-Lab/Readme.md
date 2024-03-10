### Task
You've been presented with three arrays: firstName, lastName, and age. The firstName and lastName arrays contain the first and last names of the person, respectively, while the age array stores their ages.
```
firstName : ["john","mark"]
lastName : ["Doe","wood"]
age : [23,50]

1. Index i of all the arrays(firstName, lastName, and age) represents the data of the ith person.
    --> for ex -> Index 0 contains the data of the 0th person {firstName : "John" lastName : "Doe" and age : 23}.
2. Your task is to the create an array of structs(person struct), where ith struct contains all the properties of ith person.

3.person structure.
person struct {
    firstName,
    lastName,
    age
}