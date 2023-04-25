import requests

import threading

import queue

import time

import colorama

import os

colorama.init()

def check_token(token):

    headers = {'Authorization': token}

    response = requests.get('https://discord.com/api/v8/users/@me', headers=headers)

    if response.status_code == 200:

        data = response.json()

        if 'premium_type' in data:

            if data['premium_type'] == 1:

                print(f"{Fore.GREEN}[Success] Nitro Classic on: {token}{Style.RESET_ALL}")

                with open('nitro.txt', 'a') as f:

                    f.write(token + '\n')

            elif data['premium_type'] == 2:

                print(f"{Fore.GREEN}[Success] Nitro Boost on: {token}{Style.RESET_ALL}")

                with open('nitro.txt', 'a') as f:

                    f.write(token + '\n')

            else:

                with open('invalid.txt', 'a') as f:

                    f.write(token + '\n')

                print(f"{Fore.RED}[No Nitro] {token}{Style.RESET_ALL}")

        else:

            with open('invalid.txt', 'a') as f:

                f.write(token + '\n')

            print(f"{Fore.RED}[No Nitro] {token}{Style.RESET_ALL}")

    else:

        with open('invalid.txt', 'a') as f:

            f.write(token + '\n')

        print(f"{Fore.YELLOW}[Failed] To get data: {token}{Style.RESET_ALL}")

def worker():

    while True:

        token = q.get()

        check_token(token)

        q.task_done()

try:

    num_threads = int(input("Enter number of threads to use (1-100): "))

    if num_threads < 1 or num_threads > 100:

        print("Invalid number of threads. Using default value of 10.")

        num_threads = 10

except ValueError:

    print("Invalid input. Using default value of 10.")

    num_threads = 10

q = queue.Queue()

threads = []

for i in range(num_threads):

    t = threading.Thread(target=worker)

    t.daemon = True

    t.start()

    threads.append(t)

tokens_file = 'tokens.txt'

if not os.path.exists(tokens_file):

    print(f"Error: {tokens_file} not found.")

    exit(1)

with open(tokens_file, 'r') as f:

    tokens = f.read().splitlines()

start_time = time.time()

for token in tokens:

    q.put(token)

q.join()

print(f"Finished checking {len(tokens)} tokens in {time.time() - start_time:.2f} seconds.")

