package com.example.foodapp.Activity;

import android.annotation.SuppressLint;
import android.content.Intent;
import android.content.SharedPreferences;
import android.graphics.Color;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.view.Window;
import android.view.WindowManager;

import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.foodapp.Adapter.CartAdapter;
import com.example.foodapp.Domain.Cart.CartResponse;
import com.example.foodapp.Management.CartManagement;
import com.example.foodapp.databinding.ActivityCartBinding;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class CartActivity extends BaseActivity {
    ActivityCartBinding binding;
    int userId;
    String cartId;
    CartManagement cartManagement;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        binding = ActivityCartBinding.inflate(getLayoutInflater());

        Window window = getWindow();
        window.clearFlags(WindowManager.LayoutParams.FLAG_TRANSLUCENT_STATUS);
        window.addFlags(WindowManager.LayoutParams.FLAG_DRAWS_SYSTEM_BAR_BACKGROUNDS);
        window.setStatusBarColor(Color.WHITE); // Белый цвет
        View decorView = window.getDecorView();
        decorView.setSystemUiVisibility(View.SYSTEM_UI_FLAG_LIGHT_STATUS_BAR);

        setContentView(binding.getRoot());

        cartManagement = new CartManagement(cartApi);

        SharedPreferences cartSharedPreferences = getSharedPreferences("cart_prefs", MODE_PRIVATE);
        SharedPreferences authSharedPreferences = getSharedPreferences("auth_prefs", MODE_PRIVATE);
        this.userId = authSharedPreferences.getInt("user_id", -1);
        this.cartId = cartSharedPreferences.getString("cart_id_" + userId, null);

        setVariable();

        GetCart();
    }

    private void setVariable() {
        binding.returnButtonCart.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                finish();
            }
        });
    }

    private void GetCart() {
        Call<CartResponse> call = cartApi.GetById(cartId);
        call.enqueue(new Callback<CartResponse>() {
            @SuppressLint("SetTextI18n")
            @Override
            public void onResponse(Call<CartResponse> call, Response<CartResponse> response) {
                if (response.isSuccessful() && response.body() != null) {
                    String price = response.body().getTotalPrice() + "BYN";
                    binding.priceDishesCart.setText(price);
                    binding.totalPriceOrder.setText(price);
                    binding.CartRecyclerView.setLayoutManager(new LinearLayoutManager(CartActivity.this, LinearLayoutManager.VERTICAL, false));
                    RecyclerView.Adapter<CartAdapter.viewholder> adapter = new CartAdapter(response.body().getItems(), cartManagement, cartId);
                    binding.CartRecyclerView.setAdapter(adapter);
                } else {
                    Log.d("error", response.message());
                }
            }

            @Override
            public void onFailure(Call<CartResponse> call, Throwable t) {
                Log.d("error", t.toString());
            }
        });
    }
}