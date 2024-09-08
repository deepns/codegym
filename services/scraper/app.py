import requests
from bs4 import BeautifulSoup
import time
import signal

# List of links to scrape
links = [
    "https://www.google.com",
    "https://www.apple.com",
    "https://quotes.toscrape.com",
    # Add more links as needed
]

# File to write the scraped content
output_file = "scraped_output.txt"

# Flag to control the running loop
running = True

# Function to handle signals
def handle_signal(signum, frame):
    global running
    print(f"Received signal {signum}. Shutting down gracefully...")
    running = False

# Register signal handlers
signal.signal(signal.SIGTERM, handle_signal)
signal.signal(signal.SIGINT, handle_signal)
signal.signal(signal.SIGQUIT, handle_signal)

def scrape_link(url):
    """
    Scrapes a given link, prints the list of links found and return the link contents
    """
    try:
        response = requests.get(url)
        response.raise_for_status()
        soup = BeautifulSoup(response.text, 'html.parser')

        # Find all the links on the page
        links = soup.find_all('a')

        # Extract and print the href attribute of each link
        for link in links:
            href = link.get('href')
            if href:
                print(href)

        # Extract the desired content, e.g., all text or specific tags
        content = soup.get_text()
        return content
    except requests.exceptions.RequestException as e:
        print(f"Error scraping {url}: {e}")
        return None

# Continuous scraping loop
def continuous_scrape(links, delay=60):
    while running:
        for link in links:
            if not running:
                break
            print(f"Scraping {link}")
            scrape_link(link)

        # Wait before scraping again
        time.sleep(delay)

if __name__ == "__main__":
    try:
        # Start the continuous scraping loop with a delay of 60 seconds between iterations
        continuous_scrape(links, delay=60)
    except KeyboardInterrupt:
        print("Script interrupted by user.")
    finally:
        print("Exiting...")
