Examples :

0:
Basic send and Receive. These are asynchronous i.e if a program receives a message and crashes, we might lose messages.
We use default exchange and use the routing key to route to the queue that we care about.

1:
In this example, we do work based on the number of . in the message received. The sender takes this as first optional argument.
here the messages are marked with ack (and hence they need to be ack explicitly with ch.basic_ack after processing the message
The receiver randomly exits but we notice that messages are delivered reliably.
We a received crashes and starts again (if another receiver is already working, then all pending messages are given to the other receiver queue)
and this receiver may not receive any message

2:
With basic_qos on queue and prefetch_count=1, a queue receives one message at a time. 
Hence if a receiver crashes and restarts, it still gets some messages because not all messages are delivered to other receiver during this receiver's absence.

prefetch_count is only applicable for messages with ack. When no_ack is set, prefetch_count is not applicable

3:
named exchange of type fanout. this allows multiple queues to receive the messages.
unnamed (or) exclusive queue for each process that binds to the queue

4:
direct exchange will deliver a message to a queue only if the routing_key matches to it.
if multiple queues specify the same routing key, then the message is delivered to all of them

5:
Receiver does not need to know about the exact routine key. it can take a glob - hence this concept of topic.
The publisher puts the topic (routing key) which are words seperated by .
The receiver can bin to a routing key which has either a # (representing 0 or more words) and * (which can match any word)
