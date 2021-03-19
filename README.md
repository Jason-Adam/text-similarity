# Text Similarity Service  
Microservice to detect similarity between two input strings using [Levenschtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance).  

## Run The Application  
Below are instructions for ways to run the application.  

Run the following command to pull the docker image:  

```bash  
docker pull docker.io/jasonadam/text-similarity:latest
```

Once the image has pulled, run the following to start the application:

```bash  
docker run -p 3000:3000 docker.io/jasonadam/text-similarity:latest
```  

Open a new terminal pane and you can run the following to test the endpoint with pre-defined samples:  

```bash  
curl "https://raw.githubusercontent.com/Jason-Adam/text-similarity/main/run_samples.sh" | sh
```  

Or, you can run a simple example like so:  

```bash  
curl -X \
    POST \
    "http://localhost:3000/api/v1/similarity" \
    -H  "accept: application/json" \
    -H  "Content-Type: application/json" \
    -d "{\"first_text\":\"hello world\",\"second_text\":\"hi world\"}"
```  

You should see an output like this:  

```json  
{
  "similarity_score": 0.6363636
}
```

## Testing  
The current test suite needs improvement, but it can be run with:  

```bash  
make test
```  

## Space & Time Complexity  
The Levenschtein Distance Algorithm runs in `O(nm) time | O(nm)` space, where `n` is the size of the first string and `m` is the size of the second string.
