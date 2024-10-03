package com.example.foodapp.Domain.Auth;

import com.fasterxml.jackson.annotation.JsonProperty;

public class LoginResponse {

    @JsonProperty("token")
    private String token;

    public String getToken() {
        return token;
    }
}
