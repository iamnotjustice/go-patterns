# Creational Patterns: Singleton in Go

Imagine we have some kind of config file we need to use in our program, and we need to use these settings globally.

This is where singleton comes into play: we have a single instance which reads the file upon its creation, and then we use the data where we need to just by importing a package.

So we have two requirements of a basic singleton implemented: we have a single instance (we control the number of copies) and we have a way to globally use this instance (via exported funcs).

This example shows three different ways to implement this kind of behaviour: 

- using init func;
- using mutex to lock upon first creation;
- using sync.Once to do the creation of the singleton just once.

This example uses Viper to read a config file. And this library is basically singleton, so if you want to check out a more sophisticated example - you now know where to go.