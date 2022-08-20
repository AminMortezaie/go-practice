# 1

The **fetchRemoteResource()** function we wrote simply displays the downloaded
data to the terminal. This may not serve any purpose for the application’s user
and will just appear as garbage for certain kinds of data—images and non-text
files, for example. In most cases, you may want to do some kind of processing
on the data instead.
<br/>

This processing is usually referred to as **unmarshalling**, or
**deserializing**, the data, and it involves converting bytes of data into a data structure that your application can understand. You can then perform any operation on this data structure in your application without having to query or parse the
raw bytes.
<br/>

The reverse operation is **marshalling**, or **serialization**, and it is an efficient way to convert a data structure into a data format that can then be stored or transmitted over the network.
<br/>

# 2

For example, unmarshalling bytes of
language-neutral JavaScript Object Notation (JSON) data into a slice of struct types
is a common operation. Similarly, deserializing the Go-specific **gob** (as defined
by the encoding/gob package) bytes into a struct type is another deserialization operation.
