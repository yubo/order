# Engineering Challenge Homework


## [Install GO runtime](https://golang.org/doc/install)


## Run the emulator

```
make build
make run

GO111MODULE=on go build -o order-emulator ./cmd/order-emulator
./order-emulator --filename ./var/orders.json --rate 2
emulator.go:47: emulator start, 2 order per second, total 132
emulator.go:58: Send Order 1/132 name Banana Split, temp frozen, value 0.968500, age 1s
kitchen.go:179: kitchen recv order name Banana Split, temp frozen, value 0.968500, age 1s
kitchen.go:100: place shelves[frozen] 1/10 name Banana Split, temp frozen, value 0.968500, age 1s
emulator.go:58: Send Order 2/132 name McFlury, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name McFlury, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 2/10 name McFlury, temp frozen, value 1.000000, age 0s
emulator.go:58: Send Order 3/132 name Acai Bowl, temp cold, value 0.998795, age 1s
kitchen.go:179: kitchen recv order name Acai Bowl, temp cold, value 0.998795, age 1s
kitchen.go:100: place shelves[cold] 1/10 name Acai Bowl, temp cold, value 0.998795, age 1s
emulator.go:58: Send Order 4/132 name Yogurt, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Yogurt, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 2/10 name Yogurt, temp cold, value 1.000000, age 0s
emulator.go:58: Send Order 5/132 name Chocolate Gelato, temp frozen, value 0.997967, age 1s
kitchen.go:179: kitchen recv order name Chocolate Gelato, temp frozen, value 0.997967, age 1s
kitchen.go:100: place shelves[frozen] 3/10 name Chocolate Gelato, temp frozen, value 0.997967, age 1s
emulator.go:58: Send Order 6/132 name Cobb Salad, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Cobb Salad, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 3/10 name Cobb Salad, temp cold, value 1.000000, age 0s
emulator.go:58: Send Order 7/132 name Cottage Cheese, temp cold, value 0.999124, age 1s
kitchen.go:179: kitchen recv order name Cottage Cheese, temp cold, value 0.999124, age 1s
kitchen.go:100: place shelves[cold] 4/10 name Cottage Cheese, temp cold, value 0.999124, age 1s
emulator.go:58: Send Order 8/132 name Coke, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Coke, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 5/10 name Coke, temp cold, value 1.000000, age 0s
emulator.go:80: PickUp the order name McFlury, temp frozen, value 0.996800, age 3s
emulator.go:58: Send Order 9/132 name Snow Cone, temp frozen, value 0.982800, age 1s
kitchen.go:179: kitchen recv order name Snow Cone, temp frozen, value 0.982800, age 1s
kitchen.go:100: place shelves[frozen] 3/10 name Snow Cone, temp frozen, value 0.982800, age 1s
emulator.go:58: Send Order 10/132 name Pad See Ew, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pad See Ew, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 1/10 name Pad See Ew, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Banana Split, temp frozen, value 0.842500, age 5s
emulator.go:58: Send Order 11/132 name Chunky Monkey, temp frozen, value 0.997429, age 1s
kitchen.go:179: kitchen recv order name Chunky Monkey, temp frozen, value 0.997429, age 1s
kitchen.go:100: place shelves[frozen] 3/10 name Chunky Monkey, temp frozen, value 0.997429, age 1s
emulator.go:80: PickUp the order name Chocolate Gelato, temp frozen, value 0.991867, age 4s
emulator.go:58: Send Order 12/132 name Beef Stew, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Beef Stew, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 2/10 name Beef Stew, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 13/132 name Cheese, temp cold, value 0.999216, age 1s
kitchen.go:179: kitchen recv order name Cheese, temp cold, value 0.999216, age 1s
kitchen.go:100: place shelves[cold] 6/10 name Cheese, temp cold, value 0.999216, age 1s
emulator.go:80: PickUp the order name Acai Bowl, temp cold, value 0.992771, age 6s
emulator.go:58: Send Order 14/132 name Spinach Omelet, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Spinach Omelet, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 3/10 name Spinach Omelet, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Yogurt, temp cold, value 0.992966, age 5s
emulator.go:58: Send Order 15/132 name Beef Hash, temp hot, value 0.975333, age 1s
kitchen.go:179: kitchen recv order name Beef Hash, temp hot, value 0.975333, age 1s
kitchen.go:100: place shelves[hot] 4/10 name Beef Hash, temp hot, value 0.975333, age 1s
emulator.go:80: PickUp the order name Cobb Salad, temp cold, value 0.996468, age 5s
emulator.go:80: PickUp the order name Coke, temp cold, value 0.995833, age 4s
emulator.go:80: PickUp the order name Pad See Ew, temp hot, value 0.989714, age 3s
emulator.go:80: PickUp the order name Chunky Monkey, temp frozen, value 0.992286, age 3s
emulator.go:58: Send Order 16/132 name Pork Chop, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pork Chop, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name Pork Chop, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 17/132 name Kale Salad, temp cold, value 0.999000, age 1s
kitchen.go:179: kitchen recv order name Kale Salad, temp cold, value 0.999000, age 1s
kitchen.go:100: place shelves[cold] 3/10 name Kale Salad, temp cold, value 0.999000, age 1s
emulator.go:58: Send Order 18/132 name Fresh Fruit, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Fresh Fruit, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 4/10 name Fresh Fruit, temp cold, value 1.000000, age 0s
emulator.go:80: PickUp the order name Cottage Cheese, temp cold, value 0.994741, age 6s
emulator.go:58: Send Order 19/132 name Cranberry Salad, temp cold, value 0.999143, age 1s
kitchen.go:179: kitchen recv order name Cranberry Salad, temp cold, value 0.999143, age 1s
kitchen.go:100: place shelves[cold] 4/10 name Cranberry Salad, temp cold, value 0.999143, age 1s
emulator.go:58: Send Order 20/132 name Fudge Ice Cream Cake, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Fudge Ice Cream Cake, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 2/10 name Fudge Ice Cream Cake, temp frozen, value 1.000000, age 0s
emulator.go:80: PickUp the order name Snow Cone, temp frozen, value 0.896800, age 6s
emulator.go:58: Send Order 21/132 name Mint Chocolate Ice Cream, temp frozen, value 0.998276, age 1s
kitchen.go:179: kitchen recv order name Mint Chocolate Ice Cream, temp frozen, value 0.998276, age 1s
kitchen.go:100: place shelves[frozen] 2/10 name Mint Chocolate Ice Cream, temp frozen, value 0.998276, age 1s
emulator.go:58: Send Order 22/132 name Vegan Pizza, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Vegan Pizza, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name Vegan Pizza, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 23/132 name Orange Chicken, temp hot, value 0.996884, age 1s
kitchen.go:179: kitchen recv order name Orange Chicken, temp hot, value 0.996884, age 1s
kitchen.go:100: place shelves[hot] 6/10 name Orange Chicken, temp hot, value 0.996884, age 1s
emulator.go:80: PickUp the order name Pork Chop, temp hot, value 0.986000, age 4s
emulator.go:80: PickUp the order name Beef Stew, temp hot, value 0.979903, age 6s
emulator.go:80: PickUp the order name Cranberry Salad, temp cold, value 0.997429, age 3s
emulator.go:80: PickUp the order name Spinach Omelet, temp hot, value 0.986304, age 5s
emulator.go:58: Send Order 24/132 name MeatLoaf, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name MeatLoaf, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name MeatLoaf, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Cheese, temp cold, value 0.995294, age 6s
emulator.go:58: Send Order 25/132 name Milk, temp cold, value 0.999405, age 1s
kitchen.go:179: kitchen recv order name Milk, temp cold, value 0.999405, age 1s
kitchen.go:100: place shelves[cold] 3/10 name Milk, temp cold, value 0.999405, age 1s
emulator.go:80: PickUp the order name Kale Salad, temp cold, value 0.995000, age 5s
emulator.go:80: PickUp the order name Mint Chocolate Ice Cream, temp frozen, value 0.994828, age 3s
emulator.go:58: Send Order 26/132 name Pastrami Sandwich, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pastrami Sandwich, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name Pastrami Sandwich, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Fresh Fruit, temp cold, value 0.995397, age 4s
emulator.go:80: PickUp the order name Beef Hash, temp hot, value 0.827333, age 7s
emulator.go:58: Send Order 27/132 name Arugula, temp cold, value 0.998924, age 1s
kitchen.go:179: kitchen recv order name Arugula, temp cold, value 0.998924, age 1s
kitchen.go:100: place shelves[cold] 2/10 name Arugula, temp cold, value 0.998924, age 1s
emulator.go:58: Send Order 28/132 name Pickles, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pickles, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 3/10 name Pickles, temp cold, value 1.000000, age 0s
emulator.go:58: Send Order 29/132 name Chicken, temp hot, value 0.996318, age 1s
kitchen.go:179: kitchen recv order name Chicken, temp hot, value 0.996318, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Chicken, temp hot, value 0.996318, age 1s
emulator.go:80: PickUp the order name Fudge Ice Cream Cake, temp frozen, value 0.994096, age 5s
emulator.go:58: Send Order 30/132 name Cookie Dough, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Cookie Dough, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 1/10 name Cookie Dough, temp frozen, value 1.000000, age 0s
emulator.go:80: PickUp the order name MeatLoaf, temp hot, value 0.992958, age 3s
emulator.go:58: Send Order 31/132 name Hamburger, temp hot, value 0.996850, age 1s
kitchen.go:179: kitchen recv order name Hamburger, temp hot, value 0.996850, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Hamburger, temp hot, value 0.996850, age 1s
emulator.go:80: PickUp the order name Vegan Pizza, temp hot, value 0.982500, age 5s
emulator.go:58: Send Order 32/132 name French Fries, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name French Fries, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name French Fries, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Milk, temp cold, value 0.997619, age 4s
emulator.go:58: Send Order 33/132 name Ice, temp frozen, value 0.991000, age 1s
kitchen.go:179: kitchen recv order name Ice, temp frozen, value 0.991000, age 1s
kitchen.go:100: place shelves[frozen] 2/10 name Ice, temp frozen, value 0.991000, age 1s
emulator.go:80: PickUp the order name Orange Chicken, temp hot, value 0.981302, age 6s
emulator.go:80: PickUp the order name Pastrami Sandwich, temp hot, value 0.983158, age 4s
emulator.go:58: Send Order 34/132 name Carne Asada, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Carne Asada, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name Carne Asada, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 35/132 name Sherbet, temp frozen, value 0.996571, age 1s
kitchen.go:179: kitchen recv order name Sherbet, temp frozen, value 0.996571, age 1s
kitchen.go:100: place shelves[frozen] 3/10 name Sherbet, temp frozen, value 0.996571, age 1s
emulator.go:80: PickUp the order name Cookie Dough, temp frozen, value 0.999250, age 3s
emulator.go:58: Send Order 36/132 name Orange Sorbet, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Orange Sorbet, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 3/10 name Orange Sorbet, temp frozen, value 1.000000, age 0s
emulator.go:58: Send Order 37/132 name Frosty, temp frozen, value 0.996148, age 1s
kitchen.go:179: kitchen recv order name Frosty, temp frozen, value 0.996148, age 1s
kitchen.go:100: place shelves[frozen] 4/10 name Frosty, temp frozen, value 0.996148, age 1s
emulator.go:80: PickUp the order name French Fries, temp hot, value 0.990864, age 3s
emulator.go:80: PickUp the order name Arugula, temp cold, value 0.993546, age 6s
emulator.go:58: Send Order 38/132 name Fresh Bread, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Fresh Bread, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name Fresh Bread, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Chicken, temp hot, value 0.981592, age 5s
emulator.go:80: PickUp the order name Ice, temp frozen, value 0.973000, age 3s
emulator.go:58: Send Order 39/132 name Burrito, temp hot, value 0.996436, age 1s
kitchen.go:179: kitchen recv order name Burrito, temp hot, value 0.996436, age 1s
kitchen.go:100: place shelves[hot] 4/10 name Burrito, temp hot, value 0.996436, age 1s
emulator.go:80: PickUp the order name Pickles, temp cold, value 0.993282, age 6s
emulator.go:58: Send Order 40/132 name Icy, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Icy, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 4/10 name Icy, temp frozen, value 1.000000, age 0s
emulator.go:58: Send Order 41/132 name Push Pop, temp frozen, value 0.997727, age 1s
kitchen.go:179: kitchen recv order name Push Pop, temp frozen, value 0.997727, age 1s
kitchen.go:100: place shelves[frozen] 5/10 name Push Pop, temp frozen, value 0.997727, age 1s
emulator.go:58: Send Order 42/132 name Pasta, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pasta, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name Pasta, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Hamburger, temp hot, value 0.981100, age 6s
emulator.go:58: Send Order 43/132 name Chicken Nuggets, temp hot, value 0.996537, age 1s
kitchen.go:179: kitchen recv order name Chicken Nuggets, temp hot, value 0.996537, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Chicken Nuggets, temp hot, value 0.996537, age 1s
emulator.go:80: PickUp the order name Frosty, temp frozen, value 0.984593, age 4s
emulator.go:80: PickUp the order name Burrito, temp hot, value 0.989307, age 3s
emulator.go:80: PickUp the order name Carne Asada, temp hot, value 0.984009, age 5s
emulator.go:58: Send Order 44/132 name Ice Cream Sandwich, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Ice Cream Sandwich, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 5/10 name Ice Cream Sandwich, temp frozen, value 1.000000, age 0s
emulator.go:58: Send Order 45/132 name Taco, temp hot, value 0.998081, age 1s
kitchen.go:179: kitchen recv order name Taco, temp hot, value 0.998081, age 1s
kitchen.go:100: place shelves[hot] 4/10 name Taco, temp hot, value 0.998081, age 1s
emulator.go:80: PickUp the order name Icy, temp frozen, value 0.992174, age 3s
emulator.go:58: Send Order 46/132 name Tomato Soup, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Tomato Soup, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name Tomato Soup, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Sherbet, temp frozen, value 0.979429, age 6s
emulator.go:80: PickUp the order name Orange Sorbet, temp frozen, value 0.980303, age 5s
emulator.go:58: Send Order 47/132 name Vanilla Ice Cream, temp frozen, value 0.998871, age 1s
kitchen.go:179: kitchen recv order name Vanilla Ice Cream, temp frozen, value 0.998871, age 1s
kitchen.go:100: place shelves[frozen] 3/10 name Vanilla Ice Cream, temp frozen, value 0.998871, age 1s
emulator.go:80: PickUp the order name Push Pop, temp frozen, value 0.990909, age 4s
emulator.go:80: PickUp the order name Pasta, temp hot, value 0.989500, age 3s
emulator.go:80: PickUp the order name Chicken Nuggets, temp hot, value 0.989610, age 3s
emulator.go:58: Send Order 48/132 name Poppers, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Poppers, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name Poppers, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 49/132 name Popsicle, temp frozen, value 0.997826, age 1s
kitchen.go:179: kitchen recv order name Popsicle, temp frozen, value 0.997826, age 1s
kitchen.go:100: place shelves[frozen] 3/10 name Popsicle, temp frozen, value 0.997826, age 1s
emulator.go:80: PickUp the order name Fresh Bread, temp hot, value 0.973134, age 6s
emulator.go:58: Send Order 50/132 name Strawberries, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Strawberries, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 4/10 name Strawberries, temp frozen, value 1.000000, age 0s
emulator.go:80: PickUp the order name Ice Cream Sandwich, temp frozen, value 0.994000, age 3s
emulator.go:58: Send Order 51/132 name Brown Rice, temp hot, value 0.997143, age 1s
kitchen.go:179: kitchen recv order name Brown Rice, temp hot, value 0.997143, age 1s
kitchen.go:100: place shelves[hot] 4/10 name Brown Rice, temp hot, value 0.997143, age 1s
emulator.go:58: Send Order 52/132 name Cheese Pizza, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Cheese Pizza, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name Cheese Pizza, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Taco, temp hot, value 0.992323, age 4s
emulator.go:58: Send Order 53/132 name Pressed Juice, temp cold, value 0.999200, age 1s
kitchen.go:179: kitchen recv order name Pressed Juice, temp cold, value 0.999200, age 1s
kitchen.go:100: place shelves[cold] 1/10 name Pressed Juice, temp cold, value 0.999200, age 1s
emulator.go:80: PickUp the order name Poppers, temp hot, value 0.988529, age 3s
emulator.go:80: PickUp the order name Tomato Soup, temp hot, value 0.988313, age 4s
emulator.go:58: Send Order 54/132 name Coconut, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Coconut, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 2/10 name Coconut, temp cold, value 1.000000, age 0s
emulator.go:58: Send Order 55/132 name Onion Rings, temp hot, value 0.996517, age 1s
kitchen.go:179: kitchen recv order name Onion Rings, temp hot, value 0.996517, age 1s
kitchen.go:100: place shelves[hot] 3/10 name Onion Rings, temp hot, value 0.996517, age 1s
emulator.go:58: Send Order 56/132 name Fish Tacos, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Fish Tacos, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name Fish Tacos, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Vanilla Ice Cream, temp frozen, value 0.993226, age 6s
emulator.go:58: Send Order 57/132 name Pot Stickers, temp hot, value 0.996422, age 1s
kitchen.go:179: kitchen recv order name Pot Stickers, temp hot, value 0.996422, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Pot Stickers, temp hot, value 0.996422, age 1s
emulator.go:80: PickUp the order name Strawberries, temp frozen, value 0.999600, age 4s
emulator.go:80: PickUp the order name Pressed Juice, temp cold, value 0.997600, age 3s
emulator.go:58: Send Order 58/132 name Kombucha, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Kombucha, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 2/10 name Kombucha, temp cold, value 1.000000, age 0s
emulator.go:58: Send Order 59/132 name Mixed Greens, temp cold, value 0.998968, age 1s
kitchen.go:179: kitchen recv order name Mixed Greens, temp cold, value 0.998968, age 1s
kitchen.go:100: place shelves[cold] 3/10 name Mixed Greens, temp cold, value 0.998968, age 1s
emulator.go:80: PickUp the order name Coconut, temp cold, value 0.997402, age 3s
emulator.go:58: Send Order 60/132 name Sushi, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Sushi, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 3/10 name Sushi, temp cold, value 1.000000, age 0s
emulator.go:80: PickUp the order name Onion Rings, temp hot, value 0.989552, age 3s
emulator.go:80: PickUp the order name Popsicle, temp frozen, value 0.986957, age 6s
emulator.go:58: Send Order 61/132 name Apples, temp cold, value 0.999057, age 1s
kitchen.go:179: kitchen recv order name Apples, temp cold, value 0.999057, age 1s
kitchen.go:100: place shelves[cold] 4/10 name Apples, temp cold, value 0.999057, age 1s
emulator.go:80: PickUp the order name Cheese Pizza, temp hot, value 0.981000, age 5s
emulator.go:58: Send Order 62/132 name Kebab, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Kebab, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name Kebab, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Brown Rice, temp hot, value 0.982857, age 6s
emulator.go:80: PickUp the order name Fish Tacos, temp hot, value 0.989275, age 3s
emulator.go:80: PickUp the order name Kombucha, temp cold, value 0.997683, age 3s
emulator.go:58: Send Order 63/132 name Mac & Cheese, temp hot, value 0.997512, age 1s
kitchen.go:179: kitchen recv order name Mac & Cheese, temp hot, value 0.997512, age 1s
kitchen.go:100: place shelves[hot] 3/10 name Mac & Cheese, temp hot, value 0.997512, age 1s
emulator.go:58: Send Order 64/132 name Corn Dog, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Corn Dog, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name Corn Dog, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Mixed Greens, temp cold, value 0.996905, age 3s
emulator.go:58: Send Order 65/132 name Grilled Corn Salad, temp cold, value 0.999672, age 1s
kitchen.go:179: kitchen recv order name Grilled Corn Salad, temp cold, value 0.999672, age 1s
kitchen.go:100: place shelves[cold] 3/10 name Grilled Corn Salad, temp cold, value 0.999672, age 1s
emulator.go:80: PickUp the order name Sushi, temp cold, value 0.997012, age 3s
emulator.go:80: PickUp the order name Pot Stickers, temp hot, value 0.982108, age 5s
emulator.go:58: Send Order 66/132 name Pistachio Ice Cream, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pistachio Ice Cream, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 1/10 name Pistachio Ice Cream, temp frozen, value 1.000000, age 0s
emulator.go:58: Send Order 67/132 name Strawberyy Banana Split, temp frozen, value 0.975000, age 1s
kitchen.go:179: kitchen recv order name Strawberyy Banana Split, temp frozen, value 0.975000, age 1s
kitchen.go:100: place shelves[frozen] 2/10 name Strawberyy Banana Split, temp frozen, value 0.975000, age 1s
emulator.go:58: Send Order 68/132 name McFlury, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name McFlury, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 3/10 name McFlury, temp frozen, value 1.000000, age 0s
emulator.go:58: Send Order 69/132 name Acai Bowl, temp cold, value 0.996250, age 1s
kitchen.go:179: kitchen recv order name Acai Bowl, temp cold, value 0.996250, age 1s
kitchen.go:100: place shelves[cold] 3/10 name Acai Bowl, temp cold, value 0.996250, age 1s
emulator.go:58: Send Order 70/132 name Yogurt, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Yogurt, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 4/10 name Yogurt, temp cold, value 1.000000, age 0s
emulator.go:80: PickUp the order name Grilled Corn Salad, temp cold, value 0.999016, age 3s
emulator.go:80: PickUp the order name Apples, temp cold, value 0.994344, age 6s
emulator.go:80: PickUp the order name Kebab, temp hot, value 0.986500, age 5s
emulator.go:58: Send Order 71/132 name Chocolate Gelato, temp frozen, value 0.997825, age 1s
kitchen.go:179: kitchen recv order name Chocolate Gelato, temp frozen, value 0.997825, age 1s
kitchen.go:100: place shelves[frozen] 4/10 name Chocolate Gelato, temp frozen, value 0.997825, age 1s
emulator.go:80: PickUp the order name Mac & Cheese, temp hot, value 0.987561, age 5s
emulator.go:58: Send Order 72/132 name Cobb Salad, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Cobb Salad, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 3/10 name Cobb Salad, temp cold, value 1.000000, age 0s
emulator.go:80: PickUp the order name Corn Dog, temp hot, value 0.994089, age 4s
emulator.go:58: Send Order 73/132 name Cottage Cheese, temp cold, value 0.998980, age 1s
kitchen.go:179: kitchen recv order name Cottage Cheese, temp cold, value 0.998980, age 1s
kitchen.go:100: place shelves[cold] 4/10 name Cottage Cheese, temp cold, value 0.998980, age 1s
emulator.go:58: Send Order 74/132 name Coke, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Coke, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 5/10 name Coke, temp cold, value 1.000000, age 0s
emulator.go:58: Send Order 75/132 name Snow Cone, temp frozen, value 0.982800, age 1s
kitchen.go:179: kitchen recv order name Snow Cone, temp frozen, value 0.982800, age 1s
kitchen.go:100: place shelves[frozen] 5/10 name Snow Cone, temp frozen, value 0.982800, age 1s
emulator.go:80: PickUp the order name McFlury, temp frozen, value 0.995161, age 4s
emulator.go:58: Send Order 76/132 name Pad See Ew, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pad See Ew, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 1/10 name Pad See Ew, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Pistachio Ice Cream, temp frozen, value 0.988571, age 5s
emulator.go:58: Send Order 77/132 name Chunky Monkey, temp frozen, value 0.997429, age 1s
kitchen.go:179: kitchen recv order name Chunky Monkey, temp frozen, value 0.997429, age 1s
kitchen.go:100: place shelves[frozen] 4/10 name Chunky Monkey, temp frozen, value 0.997429, age 1s
emulator.go:80: PickUp the order name Chocolate Gelato, temp frozen, value 0.991299, age 4s
emulator.go:80: PickUp the order name Yogurt, temp cold, value 0.993300, age 4s
emulator.go:58: Send Order 78/132 name Beef Stew, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Beef Stew, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 2/10 name Beef Stew, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Strawberyy Banana Split, temp frozen, value 0.850000, age 6s
emulator.go:58: Send Order 79/132 name Cheese, temp cold, value 0.999216, age 1s
kitchen.go:179: kitchen recv order name Cheese, temp cold, value 0.999216, age 1s
kitchen.go:100: place shelves[cold] 5/10 name Cheese, temp cold, value 0.999216, age 1s
emulator.go:58: Send Order 80/132 name Spinach Omelet, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Spinach Omelet, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 3/10 name Spinach Omelet, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 81/132 name Beef Hash, temp hot, value 0.975333, age 1s
kitchen.go:179: kitchen recv order name Beef Hash, temp hot, value 0.975333, age 1s
kitchen.go:100: place shelves[hot] 4/10 name Beef Hash, temp hot, value 0.975333, age 1s
emulator.go:80: PickUp the order name Acai Bowl, temp cold, value 0.973750, age 7s
emulator.go:80: PickUp the order name Cottage Cheese, temp cold, value 0.994902, age 5s
emulator.go:80: PickUp the order name Chunky Monkey, temp frozen, value 0.992286, age 3s
emulator.go:58: Send Order 82/132 name Pork Chop, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pork Chop, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name Pork Chop, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Pad See Ew, temp hot, value 0.989714, age 3s
emulator.go:80: PickUp the order name Snow Cone, temp frozen, value 0.931200, age 4s
emulator.go:80: PickUp the order name Cobb Salad, temp cold, value 0.995437, age 6s
emulator.go:58: Send Order 83/132 name Kale Salad, temp cold, value 0.999000, age 1s
kitchen.go:179: kitchen recv order name Kale Salad, temp cold, value 0.999000, age 1s
kitchen.go:100: place shelves[cold] 3/10 name Kale Salad, temp cold, value 0.999000, age 1s
emulator.go:58: Send Order 84/132 name Fresh Fruit, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Fresh Fruit, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 4/10 name Fresh Fruit, temp cold, value 1.000000, age 0s
emulator.go:80: PickUp the order name Coke, temp cold, value 0.994792, age 5s
emulator.go:58: Send Order 85/132 name Cranberry Salad, temp cold, value 0.999143, age 1s
kitchen.go:179: kitchen recv order name Cranberry Salad, temp cold, value 0.999143, age 1s
kitchen.go:100: place shelves[cold] 4/10 name Cranberry Salad, temp cold, value 0.999143, age 1s
emulator.go:80: PickUp the order name Cheese, temp cold, value 0.996863, age 4s
emulator.go:80: PickUp the order name Beef Hash, temp hot, value 0.926000, age 3s
emulator.go:58: Send Order 86/132 name Fudge Ice Cream Cake, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Fudge Ice Cream Cake, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 1/10 name Fudge Ice Cream Cake, temp frozen, value 1.000000, age 0s
emulator.go:80: PickUp the order name Beef Stew, temp hot, value 0.986602, age 4s
emulator.go:58: Send Order 87/132 name Mint Chocolate Ice Cream, temp frozen, value 0.998276, age 1s
kitchen.go:179: kitchen recv order name Mint Chocolate Ice Cream, temp frozen, value 0.998276, age 1s
kitchen.go:100: place shelves[frozen] 2/10 name Mint Chocolate Ice Cream, temp frozen, value 0.998276, age 1s
emulator.go:58: Send Order 88/132 name Vegan Pizza, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Vegan Pizza, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 3/10 name Vegan Pizza, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Fresh Fruit, temp cold, value 0.997698, age 2s
emulator.go:58: Send Order 89/132 name Orange Chicken, temp hot, value 0.996884, age 1s
kitchen.go:179: kitchen recv order name Orange Chicken, temp hot, value 0.996884, age 1s
kitchen.go:100: place shelves[hot] 4/10 name Orange Chicken, temp hot, value 0.996884, age 1s
emulator.go:80: PickUp the order name Pork Chop, temp hot, value 0.986000, age 4s
emulator.go:80: PickUp the order name Kale Salad, temp cold, value 0.996000, age 4s
emulator.go:58: Send Order 90/132 name MeatLoaf, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name MeatLoaf, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name MeatLoaf, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Fudge Ice Cream Cake, temp frozen, value 0.997639, age 2s
emulator.go:58: Send Order 91/132 name Milk, temp cold, value 0.999405, age 1s
kitchen.go:179: kitchen recv order name Milk, temp cold, value 0.999405, age 1s
kitchen.go:100: place shelves[cold] 2/10 name Milk, temp cold, value 0.999405, age 1s
emulator.go:80: PickUp the order name Spinach Omelet, temp hot, value 0.983565, age 6s
emulator.go:58: Send Order 92/132 name Pastrami Sandwich, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pastrami Sandwich, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name Pastrami Sandwich, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 93/132 name Arugula, temp cold, value 0.998924, age 1s
kitchen.go:179: kitchen recv order name Arugula, temp cold, value 0.998924, age 1s
kitchen.go:100: place shelves[cold] 3/10 name Arugula, temp cold, value 0.998924, age 1s
emulator.go:58: Send Order 94/132 name Pickles, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pickles, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 4/10 name Pickles, temp cold, value 1.000000, age 0s
emulator.go:80: PickUp the order name Mint Chocolate Ice Cream, temp frozen, value 0.993103, age 4s
emulator.go:58: Send Order 95/132 name Chicken, temp hot, value 0.996318, age 1s
kitchen.go:179: kitchen recv order name Chicken, temp hot, value 0.996318, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Chicken, temp hot, value 0.996318, age 1s
emulator.go:58: Send Order 96/132 name Cookie Dough, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Cookie Dough, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 1/10 name Cookie Dough, temp frozen, value 1.000000, age 0s
emulator.go:80: PickUp the order name Pastrami Sandwich, temp hot, value 0.991579, age 2s
emulator.go:80: PickUp the order name Cranberry Salad, temp cold, value 0.994857, age 6s
emulator.go:80: PickUp the order name Vegan Pizza, temp hot, value 0.982500, age 5s
emulator.go:58: Send Order 97/132 name Hamburger, temp hot, value 0.996850, age 1s
kitchen.go:179: kitchen recv order name Hamburger, temp hot, value 0.996850, age 1s
kitchen.go:100: place shelves[hot] 4/10 name Hamburger, temp hot, value 0.996850, age 1s
emulator.go:58: Send Order 98/132 name French Fries, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name French Fries, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name French Fries, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 99/132 name Ice, temp frozen, value 0.991000, age 1s
kitchen.go:179: kitchen recv order name Ice, temp frozen, value 0.991000, age 1s
kitchen.go:100: place shelves[frozen] 2/10 name Ice, temp frozen, value 0.991000, age 1s
emulator.go:80: PickUp the order name Milk, temp cold, value 0.997024, age 5s
emulator.go:80: PickUp the order name Orange Chicken, temp hot, value 0.981302, age 6s
emulator.go:58: Send Order 100/132 name Carne Asada, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Carne Asada, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name Carne Asada, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name MeatLoaf, temp hot, value 0.988263, age 5s
emulator.go:58: Send Order 101/132 name Sherbet, temp frozen, value 0.996571, age 1s
kitchen.go:179: kitchen recv order name Sherbet, temp frozen, value 0.996571, age 1s
kitchen.go:100: place shelves[frozen] 3/10 name Sherbet, temp frozen, value 0.996571, age 1s
emulator.go:58: Send Order 102/132 name Orange Sorbet, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Orange Sorbet, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 4/10 name Orange Sorbet, temp frozen, value 1.000000, age 0s
emulator.go:80: PickUp the order name Pickles, temp cold, value 0.995521, age 4s
emulator.go:80: PickUp the order name Cookie Dough, temp frozen, value 0.999000, age 4s
emulator.go:58: Send Order 103/132 name Frosty, temp frozen, value 0.996148, age 1s
kitchen.go:179: kitchen recv order name Frosty, temp frozen, value 0.996148, age 1s
kitchen.go:100: place shelves[frozen] 4/10 name Frosty, temp frozen, value 0.996148, age 1s
emulator.go:58: Send Order 104/132 name Fresh Bread, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Fresh Bread, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name Fresh Bread, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Arugula, temp cold, value 0.993546, age 6s
emulator.go:80: PickUp the order name Hamburger, temp hot, value 0.987400, age 4s
emulator.go:80: PickUp the order name Ice, temp frozen, value 0.973000, age 3s
emulator.go:58: Send Order 105/132 name Burrito, temp hot, value 0.996436, age 1s
kitchen.go:179: kitchen recv order name Burrito, temp hot, value 0.996436, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Burrito, temp hot, value 0.996436, age 1s
emulator.go:80: PickUp the order name Chicken, temp hot, value 0.977910, age 6s
emulator.go:58: Send Order 106/132 name Icy, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Icy, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 4/10 name Icy, temp frozen, value 1.000000, age 0s
emulator.go:80: PickUp the order name French Fries, temp hot, value 0.987818, age 4s
emulator.go:58: Send Order 107/132 name Push Pop, temp frozen, value 0.997727, age 1s
kitchen.go:179: kitchen recv order name Push Pop, temp frozen, value 0.997727, age 1s
kitchen.go:100: place shelves[frozen] 5/10 name Push Pop, temp frozen, value 0.997727, age 1s
emulator.go:80: PickUp the order name Sherbet, temp frozen, value 0.986286, age 4s
emulator.go:58: Send Order 108/132 name Pasta, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pasta, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name Pasta, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 109/132 name Chicken Nuggets, temp hot, value 0.996537, age 1s
kitchen.go:179: kitchen recv order name Chicken Nuggets, temp hot, value 0.996537, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Chicken Nuggets, temp hot, value 0.996537, age 1s
emulator.go:80: PickUp the order name Fresh Bread, temp hot, value 0.986567, age 3s
emulator.go:58: Send Order 110/132 name Ice Cream Sandwich, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Ice Cream Sandwich, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 5/10 name Ice Cream Sandwich, temp frozen, value 1.000000, age 0s
emulator.go:58: Send Order 111/132 name Taco, temp hot, value 0.998081, age 1s
kitchen.go:179: kitchen recv order name Taco, temp hot, value 0.998081, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Taco, temp hot, value 0.998081, age 1s
emulator.go:80: PickUp the order name Carne Asada, temp hot, value 0.980811, age 6s
emulator.go:58: Send Order 112/132 name Tomato Soup, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Tomato Soup, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name Tomato Soup, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 113/132 name Vanilla Ice Cream, temp frozen, value 0.998871, age 1s
kitchen.go:179: kitchen recv order name Vanilla Ice Cream, temp frozen, value 0.998871, age 1s
kitchen.go:100: place shelves[frozen] 6/10 name Vanilla Ice Cream, temp frozen, value 0.998871, age 1s
emulator.go:80: PickUp the order name Icy, temp frozen, value 0.989565, age 4s
emulator.go:80: PickUp the order name Orange Sorbet, temp frozen, value 0.976364, age 6s
emulator.go:58: Send Order 114/132 name Poppers, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Poppers, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 6/10 name Poppers, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Frosty, temp frozen, value 0.976889, age 6s
emulator.go:80: PickUp the order name Burrito, temp hot, value 0.982178, age 5s
emulator.go:58: Send Order 115/132 name Popsicle, temp frozen, value 0.997826, age 1s
kitchen.go:179: kitchen recv order name Popsicle, temp frozen, value 0.997826, age 1s
kitchen.go:100: place shelves[frozen] 4/10 name Popsicle, temp frozen, value 0.997826, age 1s
emulator.go:58: Send Order 116/132 name Strawberries, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Strawberries, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 5/10 name Strawberries, temp frozen, value 1.000000, age 0s
emulator.go:80: PickUp the order name Pasta, temp hot, value 0.986000, age 4s
emulator.go:58: Send Order 117/132 name Brown Rice, temp hot, value 0.997143, age 1s
kitchen.go:179: kitchen recv order name Brown Rice, temp hot, value 0.997143, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Brown Rice, temp hot, value 0.997143, age 1s
emulator.go:80: PickUp the order name Ice Cream Sandwich, temp frozen, value 0.992000, age 4s
emulator.go:58: Send Order 118/132 name Cheese Pizza, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Cheese Pizza, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 6/10 name Cheese Pizza, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Push Pop, temp frozen, value 0.986364, age 6s
emulator.go:58: Send Order 119/132 name Pressed Juice, temp cold, value 0.999200, age 1s
kitchen.go:179: kitchen recv order name Pressed Juice, temp cold, value 0.999200, age 1s
kitchen.go:100: place shelves[cold] 1/10 name Pressed Juice, temp cold, value 0.999200, age 1s
emulator.go:80: PickUp the order name Vanilla Ice Cream, temp frozen, value 0.995484, age 4s
emulator.go:58: Send Order 120/132 name Coconut, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Coconut, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 2/10 name Coconut, temp cold, value 1.000000, age 0s
emulator.go:80: PickUp the order name Popsicle, temp frozen, value 0.993478, age 3s
emulator.go:80: PickUp the order name Taco, temp hot, value 0.990404, age 5s
emulator.go:80: PickUp the order name Chicken Nuggets, temp hot, value 0.979220, age 6s
emulator.go:58: Send Order 121/132 name Onion Rings, temp hot, value 0.996517, age 1s
kitchen.go:179: kitchen recv order name Onion Rings, temp hot, value 0.996517, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Onion Rings, temp hot, value 0.996517, age 1s
emulator.go:80: PickUp the order name Poppers, temp hot, value 0.984706, age 4s
emulator.go:80: PickUp the order name Tomato Soup, temp hot, value 0.985391, age 5s
emulator.go:58: Send Order 122/132 name Fish Tacos, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Fish Tacos, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 4/10 name Fish Tacos, temp hot, value 1.000000, age 0s
emulator.go:58: Send Order 123/132 name Pot Stickers, temp hot, value 0.996422, age 1s
kitchen.go:179: kitchen recv order name Pot Stickers, temp hot, value 0.996422, age 1s
kitchen.go:100: place shelves[hot] 5/10 name Pot Stickers, temp hot, value 0.996422, age 1s
emulator.go:80: PickUp the order name Cheese Pizza, temp hot, value 0.988600, age 3s
emulator.go:58: Send Order 124/132 name Kombucha, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Kombucha, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 3/10 name Kombucha, temp cold, value 1.000000, age 0s
emulator.go:80: PickUp the order name Coconut, temp cold, value 0.998268, age 2s
emulator.go:80: PickUp the order name Strawberries, temp frozen, value 0.999600, age 4s
emulator.go:58: Send Order 125/132 name Mixed Greens, temp cold, value 0.998968, age 1s
kitchen.go:179: kitchen recv order name Mixed Greens, temp cold, value 0.998968, age 1s
kitchen.go:100: place shelves[cold] 3/10 name Mixed Greens, temp cold, value 0.998968, age 1s
emulator.go:58: Send Order 126/132 name Sushi, temp cold, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Sushi, temp cold, value 1.000000, age 0s
kitchen.go:100: place shelves[cold] 4/10 name Sushi, temp cold, value 1.000000, age 0s
emulator.go:58: Send Order 127/132 name Apples, temp cold, value 0.999057, age 1s
kitchen.go:179: kitchen recv order name Apples, temp cold, value 0.999057, age 1s
kitchen.go:100: place shelves[cold] 5/10 name Apples, temp cold, value 0.999057, age 1s
emulator.go:58: Send Order 128/132 name Kebab, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Kebab, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 5/10 name Kebab, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Kombucha, temp cold, value 0.998455, age 2s
emulator.go:80: PickUp the order name Brown Rice, temp hot, value 0.982857, age 6s
emulator.go:80: PickUp the order name Pot Stickers, temp hot, value 0.989265, age 3s
emulator.go:58: Send Order 129/132 name Mac & Cheese, temp hot, value 0.997512, age 1s
kitchen.go:179: kitchen recv order name Mac & Cheese, temp hot, value 0.997512, age 1s
kitchen.go:100: place shelves[hot] 4/10 name Mac & Cheese, temp hot, value 0.997512, age 1s
emulator.go:80: PickUp the order name Onion Rings, temp hot, value 0.982587, age 5s
emulator.go:80: PickUp the order name Fish Tacos, temp hot, value 0.985700, age 4s
emulator.go:58: Send Order 130/132 name Corn Dog, temp hot, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Corn Dog, temp hot, value 1.000000, age 0s
kitchen.go:100: place shelves[hot] 3/10 name Corn Dog, temp hot, value 1.000000, age 0s
emulator.go:80: PickUp the order name Pressed Juice, temp cold, value 0.995200, age 6s
emulator.go:58: Send Order 131/132 name Grilled Corn Salad, temp cold, value 0.999672, age 1s
kitchen.go:179: kitchen recv order name Grilled Corn Salad, temp cold, value 0.999672, age 1s
kitchen.go:100: place shelves[cold] 4/10 name Grilled Corn Salad, temp cold, value 0.999672, age 1s
emulator.go:80: PickUp the order name Sushi, temp cold, value 0.997012, age 3s
emulator.go:80: PickUp the order name Apples, temp cold, value 0.997172, age 3s
emulator.go:58: Send Order 132/132 name Pistachio Ice Cream, temp frozen, value 1.000000, age 0s
kitchen.go:179: kitchen recv order name Pistachio Ice Cream, temp frozen, value 1.000000, age 0s
kitchen.go:100: place shelves[frozen] 1/10 name Pistachio Ice Cream, temp frozen, value 1.000000, age 0s
emulator.go:80: PickUp the order name Mixed Greens, temp cold, value 0.995873, age 4s
emulator.go:80: PickUp the order name Kebab, temp hot, value 0.994600, age 2s
emulator.go:80: PickUp the order name Mac & Cheese, temp hot, value 0.992537, age 3s
emulator.go:80: PickUp the order name Pistachio Ice Cream, temp frozen, value 0.993143, age 3s
emulator.go:80: PickUp the order name Corn Dog, temp hot, value 0.994089, age 4s
emulator.go:80: PickUp the order name Grilled Corn Salad, temp cold, value 0.998033, age 6s


```
