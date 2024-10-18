package com.example.foodapp.Domain.Cart;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CartCreateResponse {
    @JsonProperty("id")
    String id;

    public String getId() {
        return id;
    }
}
