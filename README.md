**Epsilon** is a 64-bit ID generator similar to [Twitter Snowflake](https://github.com/twitter-archive/snowflake).

- ID is composed of:
  - Timestamp 45bits
    - increases in units of about 100µs(microseconds).
    - can be stored for a total of 68 years.
  - ParentsID 9bits
  - SequenceNumber 10bits
