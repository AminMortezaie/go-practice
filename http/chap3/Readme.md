# 1

The fetchRemoteResource() function we wrote simply displays the downloaded
data to the terminal. This may not serve any purpose for the application’s user
and will just appear as garbage for certain kinds of data—images and non-text
files, for example. In most cases, you may want to do some kind of processing
on the data instead.
<br/>

This processing is usually referred to as **unmarshalling**, or
**deserializing**, the data, and it involves converting bytes of data into a data structure that your application can understand.
