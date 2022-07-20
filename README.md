# Mastering Multithreading Programming with Go
The mood in the meeting on the 12th floor of an international investment bank was as bleak as it gets. The developers of the firm met to discuss the best way forward after a critical core application failed and caused a system wide outage.



"Guys, we have a serious issue here. I found out that the outage was caused by a race condition in our code, introduced a while ago and triggered last night." says Mark Adams, senior developer.



The room goes silent. The cars outside the floor to ceiling windows slowly and silently creep along in the heavy city traffic. The senior developers immediately understand the severity of the situation, realizing that they will now be working around the clock to fix the issue and sort out the mess in the datastore. The less experienced developers understand that a race condition is serious but don't know exactly what causes it and therefore keep their mouths shut.



Eventually Brian Holmes, delivery manager, breaks the silence with "The application has been running for months without any problems, we haven't released any code recently, how is it possible that the software just broke down?!"



Everyone shakes their heads and goes back to their desk leaving Brian in the room alone, puzzled. He takes out his phone and googles "race condition".



Sound familiar? How many times have you heard another developer talking about using threads and concurrent programming to solve a particular problem but out of fear you stayed out of the discussion?



Here's the little secret that senior developers will never share... Multithreading programming is not much harder than normal programming. Developers are scared of concurrent programming because they think it is an advanced topic that only highly experienced developers get to play with.



This is far from the truth. Our minds are very much used to dealing with concurrency. In fact we do this in our everyday life without any problem but somehow we struggle to translate this into our code. One of the reasons for this is that we're not familiar with the concepts and tools available to us to manage this concurrency. This course is here to help you understand how to use multithreading tools and concepts to manage your parallel programming. It is designed to be as practical as possible. We start with some theory around parallelism and then explain how the operating system handles multiple processes and threads. Later we move on to explain the multiple tools available by solving example problems using multithreading.



In this course we use Google's Go programming language with its goroutines, however the concepts learned here can be applied to most programming languages.

**All code in this course can be found on github, username/project: cutajarj/multithreadingingo**
