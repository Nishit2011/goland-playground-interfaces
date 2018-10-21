Why are interfaces used?

Every value that we declare and program has to have a type assigned to it.
Even if we don't explicitly define the type, := syntax figures out the type for us. So, every value still truly has a type.

Another important thing is that every function that we have written is always specifying the type of its arguments and of its receiver.

If we look at golang-playground(project on github), we saw that our shuffle function has a type of deck.
So, every function has a type. 

So, does that mean that every function that is every written in Go have to be re written to handle some 
different type even if the logic inside remains the same?????

So, does that mean that if our shuffle function deck type ,we write one function and when it deals with some other type, we repeat the function by just changing the type passed???

No.

Here come the role of interfaces.