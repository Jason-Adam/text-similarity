# Text Similarity Service  
Microservice to detect similarity between two input strings using [Levenschtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance).  

## Run The Application  
Below are instructions for ways to run the application.  

### Locally (no Docker)  
Run the following command:  

```bash  
make run
```

This will start the application. You can open another terminal window and run the following:  

```bash  
sh run_samples.sh
```  

You should see an output similar to this:  

```bash  
Sample 1: The easiest way to earn points with Fetch Rewards is to just shop for the products you already love. If you have any participating brands on your receipt, you'll get points based on the cost of the products. You don't need to clip any coupons or scan individual barcodes. Just scan each grocery receipt after you shop and we'll find the savings for you.

-----------------------------------------
Sample 2: The easiest way to earn points with Fetch Rewards is to just shop for the items you already buy. If you have any eligible brands on your receipt, you will get points based on the total cost of the products. You do not need to cut out any coupons or scan individual UPCs. Just scan your receipt after you check out and we will find the savings for you.

-----------------------------------------
Sample 3: We are always looking for opportunities for you to earn more points, which is why we also give you a selection of Special Offers. These Special Offers are opportunities to earn bonus points on top of the regular points you earn every time you purchase a participating brand. No need to pre-select these offers, we'll give you the points whether or not you knew about the offer. We just think it is easier that way.

-----------------------------------------

Comparing Sample 1 and Sample 2
{"similarity_score":0.8130312}
--------------------------------
Comparing Sample 1 and Sample 3
{"similarity_score":0.2801932}
--------------------------------
Comparing Sample 2 and Sample 3
{"similarity_score":0.29710144}
```
