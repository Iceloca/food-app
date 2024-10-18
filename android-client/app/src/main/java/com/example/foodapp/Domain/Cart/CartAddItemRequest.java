package com.example.foodapp.Domain.Cart;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CartAddItemRequest {
    @JsonProperty("basket_id")
    private String basketId;

    @JsonProperty("item")
    private Item item;

    public CartAddItemRequest(Item item, String basketId) {
        this.item = item;
        this.basketId = basketId;
    }
}
