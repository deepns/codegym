from tqdm import tqdm
import time


for i in tqdm(range(100), desc="Progress", unit="ops"):
    time.sleep(0.02)


for i in tqdm(range(100), desc="Loading", unit="files", colour="green", bar_format="{l_bar}{bar:20}{r_bar}{bar:-20b}"):
    time.sleep(0.02)

for i in tqdm(range(5), desc="Outer Loop", bar_format="{l_bar}{bar:20}{r_bar}{bar:-20b}"):
    for j in tqdm(range(50), desc="Inner Loop", leave=False, unit="item"):
        time.sleep(0.01)

for i in tqdm(range(100), desc="Downloading", ascii="▇▁", bar_format="{l_bar}{bar:20}{r_bar}{bar:-20b}"):
    time.sleep(0.02)

pbar = tqdm(total=100, desc="Processing", unit="task")
for i in range(10):
    time.sleep(0.1)
    pbar.update(10)
pbar.close()

with tqdm(total=100, desc="Transferring", unit="MB", miniters=1, dynamic_ncols=True) as pbar:
    for i in range(10):
        time.sleep(0.2)
        pbar.update(10)