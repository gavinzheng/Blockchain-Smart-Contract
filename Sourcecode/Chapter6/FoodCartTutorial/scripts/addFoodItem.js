let addFoodItemToCart = (foodName, price) => {
    foodCart.addFoodItem(foodName, price)
        .then((trxn) => {
            const details = trxn.logs[0].args;
            details.sku = details.sku.toNumber();
            details.price = details.price.toNumber();
            details.state = trxn.logs[0].event;
            console.log(details);
        });
}