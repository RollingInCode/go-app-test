# Go Data Pipeline sample

This pipeline reads one record at a time from the CSV file, converts it to a User object, then encodes it in JSON and writes it to the JSON file. All the reading, processing, and writing is done concurrently, which makes it very efficient, even for large datasets.