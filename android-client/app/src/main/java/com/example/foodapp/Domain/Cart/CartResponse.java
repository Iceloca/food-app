package com.example.foodapp.Domain.Cart;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.ArrayList;

public class CartResponse {
    @JsonProperty("id")
    private String basketId;

    @JsonProperty("user_id")
    private int userId;

    @JsonProperty("items")
    private ArrayList<Item> items;

    @JsonProperty("total_price")
    private float totalPrice;

    public String getBasketId() {
        return basketId;
    }

    public float getTotalPrice() {
        return totalPrice;
    }

    public ArrayList<Item> getItems() {
        return items;
    }

    public int getUserId() {
        return userId;
    }
}
