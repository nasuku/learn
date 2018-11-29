#!/usr/bin/env python3
import pika
import sys

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()
channel.queue_declare(queue='hello2',durable=True)


message = ' '.join(sys.argv[1:]) or "Hello World!"

channel.basic_publish(exchange='',
                      routing_key='hello2',
                      body=message,
                      properties=pika.BasicProperties(
                         delivery_mode = 2, # make message persistent
                      ) )
print(" [x] Sent %r " % message)
connection.close()
