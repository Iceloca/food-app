package com.example.foodapp.Domain.Auth;

import com.fasterxml.jackson.annotation.JsonProperty;

public class RegisterResponse {

    @JsonProperty("user_id")
    private int id;

    public int getId() {
        return this.id;
    }
}
