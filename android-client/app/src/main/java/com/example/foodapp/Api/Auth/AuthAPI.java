package com.example.foodapp.Api.Auth;


import com.example.foodapp.Domain.Auth.LoginRequest;
import com.example.foodapp.Domain.Auth.LoginResponse;
import com.example.foodapp.Domain.Auth.RegisterRequest;
import com.example.foodapp.Domain.Auth.RegisterResponse;

import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.POST;

public interface AuthAPI {
    @POST("api/v1/auth/login")
    Call<LoginResponse> Login(@Body LoginRequest loginRequest);
    @POST("api/v1/auth/register")
    Call<RegisterResponse> Register(@Body RegisterRequest registerRequest);
}


