package com.example.foodapp.Domain.Auth;

import com.fasterxml.jackson.annotation.JsonProperty;

public class RegisterRequest {

    @JsonProperty("email")
    private final String email;

    @JsonProperty("password")
    private final String password;

    public RegisterRequest(String email, String password) {
        this.password = password;
        this.email = email;
    }
}
