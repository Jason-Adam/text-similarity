#!/bin/sh

set -e

SAMPLE_ONE="The easiest way to earn points with Fetch Rewards is to just shop for the products you already love. If you have any participating brands on your receipt, you'll get points based on the cost of the products. You don't need to clip any coupons or scan individual barcodes. Just scan each grocery receipt after you shop and we'll find the savings for you."
SAMPLE_TWO="The easiest way to earn points with Fetch Rewards is to just shop for the items you already buy. If you have any eligible brands on your receipt, you will get points based on the total cost of the products. You do not need to cut out any coupons or scan individual UPCs. Just scan your receipt after you check out and we will find the savings for you."
SAMPLE_THREE="We are always looking for opportunities for you to earn more points, which is why we also give you a selection of Special Offers. These Special Offers are opportunities to earn bonus points on top of the regular points you earn every time you purchase a participating brand. No need to pre-select these offers, we'll give you the points whether or not you knew about the offer. We just think it is easier that way."

echo "Sample 1: $SAMPLE_ONE"
echo ""
echo "-----------------------------------------"

echo "Sample 2: $SAMPLE_TWO"
echo ""
echo "-----------------------------------------"

echo "Sample 3: $SAMPLE_THREE"
echo ""
echo "-----------------------------------------"
echo ""

echo "Comparing Sample 1 and Sample 2"

curl -X \
    POST "http://localhost:3000/api/v1/similarity" \
    -H "accept: application/json" \
    -H "Content-Type: application/json" \
    -d "{\"first_text\":\"$SAMPLE_ONE\",\"second_text\":\"$SAMPLE_TWO\"}"

echo ""
echo "--------------------------------"
echo "Comparing Sample 1 and Sample 3"

curl -X \
    POST "http://localhost:3000/api/v1/similarity" \
    -H "accept: application/json" \
    -H "Content-Type: application/json" \
    -d "{\"first_text\":\"$SAMPLE_ONE\",\"second_text\":\"$SAMPLE_THREE\"}"

echo ""
echo "--------------------------------"
echo "Comparing Sample 2 and Sample 3"

curl -X \
    POST "http://localhost:3000/api/v1/similarity" \
    -H "accept: application/json" \
    -H "Content-Type: application/json" \
    -d "{\"first_text\":\"$SAMPLE_TWO\",\"second_text\":\"$SAMPLE_THREE\"}"
