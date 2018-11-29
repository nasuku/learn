#!/usr/bin/env python3
import pika
import time
import random
import sys

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()
channel.queue_declare(queue='hello2',durable=True)

def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)
    time.sleep(body.count(b'.'))
    if random.choice(range(0,10)) == 0:
        print("Exiting randomly")
        sys.exit(1)
    print(" [x] Done")
    ch.basic_ack(delivery_tag = method.delivery_tag)


channel.basic_qos(prefetch_count=1) # give messages one after other
channel.basic_consume(callback,
                      queue='hello2')

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()
