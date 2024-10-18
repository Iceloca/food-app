package com.example.foodapp.Activity;

import android.os.Bundle;

import androidx.appcompat.app.AppCompatActivity;

import com.example.foodapp.Api.Cart.CartAPI;
import com.example.foodapp.R;
import com.example.foodapp.Api.Auth.AuthAPI;
import com.example.foodapp.Api.Product.ProductAPI;

import java.util.concurrent.TimeUnit;

import okhttp3.OkHttpClient;
import retrofit2.Retrofit;
import retrofit2.converter.jackson.JacksonConverterFactory;

public class BaseActivity extends AppCompatActivity {

    protected AuthAPI authApi;
    protected ProductAPI productApi;
    protected CartAPI cartApi;

    public static String errorConnection = "Ошибка соединения";
    public static String incorrectEmail = "Неккоректный адрес электронной почты";
    public static String incorrectPassword = "Некорректный пароль. Пароль должен содержать только английский буквы, цифры, спец. символы";

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        Retrofit retrofit = new Retrofit.Builder()
                .baseUrl("http://192.168.100.68:8080/")
                .addConverterFactory(JacksonConverterFactory.create())
               .build();

        authApi = retrofit.create(AuthAPI.class);
        productApi = retrofit.create(ProductAPI.class);
        cartApi = retrofit.create(CartAPI.class);
        getWindow().setStatusBarColor(getResources().getColor(R.color.white));
    }
}


