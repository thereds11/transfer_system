import json
import logging
import os
# from kafka import KafkaConsumer
import psycopg2
from aiokafka import AIOKafkaConsumer, AIOKafkaProducer
import asyncio

# Setup logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s [%(levelname)s] %(message)s')

# Load env variables
KAFKA_BROKER = os.getenv("KAFKA_BROKER", "localhost:9092")
KAFKA_TOPIC = os.getenv("KAFKA_TOPIC", "transaction.created")
FRAUD_TOPIC = os.getenv("FRAUD_TOPIC", "transaction.fraud")
DB_URL = os.getenv("DB_URL", "dbname=transferdb user=transferuser password=secret host=localhost")

# Connect to Postgres
def get_db_connection():
    try:
        conn = psycopg2.connect(DB_URL)
        conn.autocommit = True
        return conn
    except Exception as e:
        logging.error(f"Failed to connect to DB: {e}")
        exit(1)

# Process fraud detection
def is_fraudulent(tx):
    return tx.get("amount", 0) > 5000

# Update transaction as fraudulent in DB
def flag_transaction_as_fraud(conn, source_id):
    try:
        with conn.cursor() as cur:
            cur.execute("""
                UPDATE transactions
                SET fraud_flag = true
                WHERE source_account_id = %s
            """, (source_id,))
            logging.info(f"🚨 Transaction flagged as fraudulent for account: {source_id}")
    except Exception as e:
        logging.error(f"DB update failed: {e}")

async def emit_fraud_event(producer, tx):
    try:
        await producer.send_and_wait(FRAUD_TOPIC, tx)
        logging.info(f"📤 Emitted fraud event to topic '{FRAUD_TOPIC}': {tx}")
    except Exception as e:
        logging.error(f"❌ Failed to emit fraud event: {e}")
    

async def consume():
    logging.info(f"🚀 Starting Fraud Detection Service on topic: {KAFKA_TOPIC}")
    logging.info(KAFKA_BROKER)
    conn = get_db_connection()
    consumer = AIOKafkaConsumer(
        KAFKA_TOPIC,
        bootstrap_servers=[KAFKA_BROKER],
        group_id="fraud-consumer-1",
        value_deserializer=lambda x: json.loads(x.decode('utf-8')),
    )
    producer = AIOKafkaProducer(
        bootstrap_servers=[KAFKA_BROKER],
        value_serializer=lambda v: json.dumps(v).encode('utf-8')
    )
    await consumer.start()
    await producer.start()
    try:
        async for message in consumer:
            logging.info(f"📨 Message received (partition {message.partition}, offset {message.offset})")
            tx = message.value
            logging.info(f"🔍 Processing transaction: {tx}")
            if is_fraudulent(tx):
                await emit_fraud_event(producer, tx)
            else:
                logging.info("✅ Transaction passed fraud check")
    finally:
        await consumer.stop()
        await producer.stop()
asyncio.run(consume())