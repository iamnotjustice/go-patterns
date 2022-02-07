# Structural patterns in Go: Adapter

In this example we make an adapter which helps us traslate big-endian bytes into little-endian. 

Let's imagine we have our machine-code for PowerPC architecture, which uses BigEndian byte order. But for some reason we need to run this code on x86 machine. For simplicity here we assume that the byte order is the *only* difference between two execution types.

From the code perspective they not only have different interfaces (PPC runner and x86 one), but they also use different byte-order. Both of the differences we need to handle in our adapter.

We let's adapt!

We create an adapter struct which converts different parts of execution command from PPC to x86, but the calling code does not even know that the code is being "run" on x86, not PPC. So we (at least partially) covered some of the complexity.