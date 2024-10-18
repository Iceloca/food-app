package com.example.foodapp.Domain.Cart;

import com.example.foodapp.Domain.Product.ProductResponse;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

public class Item {
    @JsonProperty("product")
    private ProductResponse product;

    @JsonProperty("count")
    private int count;

    public Item(
            @JsonProperty("product") ProductResponse product,
            @JsonProperty("count") int count) {
        this.product = product;
        this.count = count;
    }

    public void setCount(int count) {
        this.count = count;
    }

    public ProductResponse getProduct() {
        return product;
    }

    public int getCount() {
        return count;
    }
}
