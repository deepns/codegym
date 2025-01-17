import os
import time
from pymongo import MongoClient

class SnowflakeIDGenerator:
    def __init__(self, epoch: int = 1640995200000):
        self.machine_id = int(os.getenv("MACHINE_ID", "0")) & 0x3FF
        self.epoch = epoch
        self.sequence = 0
        self.last_timestamp = -1

    def _current_timestamp(self):
        return int(time.time() * 1000)

    def generate_id(self):
        timestamp = self._current_timestamp()

        if timestamp == self.last_timestamp:
            self.sequence = (self.sequence + 1) & 0xFFF
            if self.sequence == 0:
                while timestamp <= self.last_timestamp:
                    timestamp = self._current_timestamp()
        else:
            self.sequence = 0

        self.last_timestamp = timestamp

        return (
            ((timestamp - self.epoch) << 22) |
            (self.machine_id << 12) |
            self.sequence
        )

# MongoDB connection
client = MongoClient(os.getenv("MONGO_URI", "mongodb://mongo:27017/"))
db = client["test_database"]
collection = db["test_collection"]

# Initialize Snowflake Generator
generator = SnowflakeIDGenerator()

# Generate and insert IDs
batch_size = int(os.getenv("BATCH_SIZE", "10000"))
for i in range(batch_size):
    unique_id = generator.generate_id()
    document = {
        "_id": unique_id,
        "pod": os.getenv("HOSTNAME", "unknown"),
        "timestamp": time.time()
    }
    collection.insert_one(document)

print(f"Generated and inserted {batch_size} IDs.")
