import asyncio


# Define a coroutine that simulates a time-consuming task
async def fetch_data(delay, id):
    print(f"Fetching data... id: {id}")
    await asyncio.sleep(delay)
    print(f"Data fetched... id: {id}")
    return {
        "data": "Some data",
        "id": id
    }


# Define another couroutine that calls the first coroutine
async def main():
    print("Start of main coroutine")
    
    task1 = fetch_data(2, 1)
    result1 = await task1
    print(f"Received result: {result1}")

    task2 = fetch_data(2, 2)
    result2 = await task2
    print(f"Received result: {result2}")

    task3 = fetch_data(2, 3)
    result3 = await task3
    print(f"Received result: {result3}")

    print("End of main coroutine")


asyncio.run(main())
