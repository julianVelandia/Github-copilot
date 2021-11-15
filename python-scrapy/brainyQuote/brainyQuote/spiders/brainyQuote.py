import scrapy

url = 'https://www.brainyquote.com/topics/scrape-quotes'
quoteXpath = '//a[contains(@class, "b-qt")]//text()'
authorXpath = '//a[contains(@class, "bq-aut")]//text()'

class SpiderBrainyQuote(scrapy.Spider):
    name = 'brainyQuote'
    start_urls = [url]

    custom_settings = {
        'FEED_FORMAT': 'json',
        'FEED_URI': 'brainyQuote.json',
        'FEED_EXPORT_FIELDS': ['quote', 'author'],
        'FEED_EXPORT_ENCODING': 'utf-8'
    }

    def parse(self, response):
        for quote in response.xpath(quoteXpath):
            if quote.get() != '\n':
                yield {
                    'quote': quote.get(),
                    'author': response.xpath(authorXpath).get()
                }