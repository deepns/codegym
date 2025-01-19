import os
import time
from pymongo import MongoClient
from pymongo.errors import DuplicateKeyError

class SnowflakeIDGenerator:
    """
    A class to generate unique Snowflake IDs.

    Attributes:
        epoch (int): The custom epoch timestamp in milliseconds. Default is 1640995200000.
        machine_id (int): The machine ID, derived from the environment variable "MACHINE_ID" and masked with 0x3FF.
        sequence (int): The sequence number for IDs generated within the same millisecond.
        last_timestamp (int): The timestamp of the last generated ID.

    Methods:
        _current_timestamp(): Returns the current timestamp in milliseconds.
        generate_id(): Generates a unique Snowflake ID based on the current timestamp, machine ID, and sequence number.
    """
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
    try:
        collection.insert_one(document)
    except DuplicateKeyError:
        print(f"Document with _id {document['_id']} already exists.")

print(f"Generated and inserted {batch_size} IDs.")
