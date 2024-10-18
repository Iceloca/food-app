package com.example.foodapp.Domain.Cart;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CartAddItemResponse {
    @JsonProperty("success")
    boolean success;
}
