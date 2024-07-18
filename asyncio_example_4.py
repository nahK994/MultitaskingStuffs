import asyncio
from datetime import datetime as dt


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
    
    task1 = asyncio.create_task(fetch_data(3, 1))
    task2 = asyncio.create_task(fetch_data(5, 2))
    task3 = asyncio.create_task(fetch_data(2, 3))
    task4 = asyncio.create_task(fetch_data(1, 4))
    task5 = asyncio.create_task(fetch_data(4, 5))

    results = await asyncio.gather(task1, task2, task3, task4, task5)
    for result in results:
        print(result)

    print("End of main coroutine")


t1 = dt.now()
asyncio.run(main())
t2 = dt.now()
print(f"Taken time {t2-t1}")
