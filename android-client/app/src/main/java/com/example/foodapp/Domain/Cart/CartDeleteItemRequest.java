package com.example.foodapp.Domain.Cart;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CartDeleteItemRequest {
    @JsonProperty("basket_id")
    private String basketId;

    @JsonProperty("item_id")
    private int itemId;

    public CartDeleteItemRequest(String basketId, int itemId) {
        this.basketId = basketId;
        this.itemId = itemId;
    }
}
