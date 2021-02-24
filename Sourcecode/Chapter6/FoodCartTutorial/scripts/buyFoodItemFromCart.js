let buyFoodItemFromCart = (itemSku, amount) => {
    foodCart.buyFoodItem(itemSku, {from: buyerAddress, value: amount})
        .then((trxn) => {
            const details = trxn.logs[0].args;
            details.sku = details.sku.toNumber();
            details.price = details.price.toNumber();
            details.state = trxn.logs[0].event;
            console.log(details);
        });
}