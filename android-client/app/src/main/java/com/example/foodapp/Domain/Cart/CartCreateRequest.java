package com.example.foodapp.Domain.Cart;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CartCreateRequest {
    @JsonProperty("user_id")
    private int userId;

    public CartCreateRequest(int userId) {
        this.userId = userId;
    }
}
