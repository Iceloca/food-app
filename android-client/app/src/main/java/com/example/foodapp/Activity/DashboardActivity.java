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

import androidx.recyclerview.widget.GridLayoutManager;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.foodapp.Adapter.BestProductAdapter;
import com.example.foodapp.Adapter.CategoryAdapter;
import com.example.foodapp.Domain.Cart.CartCreateRequest;
import com.example.foodapp.Domain.Cart.CartCreateResponse;
import com.example.foodapp.Domain.Product.CategoryResponse;
import com.example.foodapp.Domain.Product.ProductResponse;
import com.example.foodapp.databinding.ActivityDashboardBinding;

import java.util.ArrayList;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class DashboardActivity extends BaseActivity {
    ActivityDashboardBinding binding;
    String cartId;
    int userId;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        binding = ActivityDashboardBinding.inflate(getLayoutInflater());

        Window window = getWindow();
        window.clearFlags(WindowManager.LayoutParams.FLAG_TRANSLUCENT_STATUS);
        window.addFlags(WindowManager.LayoutParams.FLAG_DRAWS_SYSTEM_BAR_BACKGROUNDS);
        window.setStatusBarColor(Color.WHITE); // Белый цвет
        View decorView = window.getDecorView();
        decorView.setSystemUiVisibility(View.SYSTEM_UI_FLAG_LIGHT_STATUS_BAR);

        setContentView(binding.getRoot());

        SharedPreferences cartSharedPreferences = getSharedPreferences("cart_prefs", MODE_PRIVATE);
        SharedPreferences authSharedPreferences = getSharedPreferences("auth_prefs", MODE_PRIVATE);
        this.userId = authSharedPreferences.getInt("user_id", -1);
        this.cartId = cartSharedPreferences.getString("cart_id_" + userId, null);

        setVariable();
        GetBestProducts();
    }

    private void setVariable() {
        binding.logoutButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                SharedPreferences sharedPreferences = getSharedPreferences("auth_prefs", MODE_PRIVATE);
                SharedPreferences.Editor editor = sharedPreferences.edit();
                editor.remove("jwt_token");
                editor.apply();
                startActivity(new Intent(DashboardActivity.this, LoginActivity.class));
                finish();
            }
        });

        binding.cartButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                if (cartId == null) {
                    createCart();
                }
                Intent intent = new Intent(DashboardActivity.this, CartActivity.class);

                startActivity(intent);
            }
        });
    }


    private void GetBestProducts() {
        Call<ArrayList<ProductResponse>> call = productApi.GetBestProducts();
        binding.progressBarBestFood.setVisibility(View.VISIBLE);
        call.enqueue(new Callback<ArrayList<ProductResponse>>() {
            @Override
            public void onResponse(Call<ArrayList<ProductResponse>> call, Response<ArrayList<ProductResponse>> response) {
                if (response.isSuccessful() && response.body() != null) {
                    binding.bestProductRecycleView.setLayoutManager(
                            new LinearLayoutManager(DashboardActivity.this,
                                    LinearLayoutManager.HORIZONTAL,
                                    false));
                    RecyclerView.Adapter<BestProductAdapter.viewholder> adapter = new BestProductAdapter(response.body());
                    binding.bestProductRecycleView.setAdapter(adapter);
                    binding.progressBarBestFood.setVisibility(View.GONE);
                    GetCategories();
                } else {
                    Log.d("error", response.message());
                }
            }

            @Override
            public void onFailure(Call<ArrayList<ProductResponse>> call, Throwable t) {
                Log.d("error", t.toString());
            }
        });
    }

    private void GetCategories() {
        Call<ArrayList<CategoryResponse>> call = productApi.GetCategories();
        binding.progressBarCategory.setVisibility(View.VISIBLE);
        call.enqueue(new Callback<ArrayList<CategoryResponse>>() {
            @Override
            public void onResponse(Call<ArrayList<CategoryResponse>> call, Response<ArrayList<CategoryResponse>> response) {
                if (response.isSuccessful() && response.body() != null) {
                    binding.categoryRecycleView.setLayoutManager(new GridLayoutManager(DashboardActivity.this, 4));
                    RecyclerView.Adapter<CategoryAdapter.viewholder> adapter = new CategoryAdapter(response.body());
                    binding.categoryRecycleView.setAdapter(adapter);
                    binding.progressBarCategory.setVisibility(View.GONE);
                } else {
                    Log.d("error", response.message());
                }
            }

            @Override
            public void onFailure(Call<ArrayList<CategoryResponse>> call, Throwable t) {
                Log.d("error", t.toString());
            }
        });
    }

    private void createCart() {
        SharedPreferences sharedPreferences = getSharedPreferences("auth_prefs", MODE_PRIVATE);
        int userId = sharedPreferences.getInt("user_id",  -1);
        CartCreateRequest req = new CartCreateRequest(userId);
        Call<CartCreateResponse> call = cartApi.Create(req);

        call.enqueue(new Callback<CartCreateResponse>() {
            @Override
            public void onResponse(Call<CartCreateResponse> call, Response<CartCreateResponse> response) {
                if (response.isSuccessful() && response.body() != null) {
                    SharedPreferences sharedPreferences = getSharedPreferences("cart_prefs", MODE_PRIVATE);
                    @SuppressLint("CommitPrefEdits") SharedPreferences.Editor editor = sharedPreferences.edit();
                    cartId = response.body().getId();
                    editor.putString("cart_id_" + userId, cartId);
                    editor.apply();
                } else {
                    Log.d("error", response.message());
                }
            }

            @Override
            public void onFailure(Call<CartCreateResponse> call, Throwable t) {
                Log.d("error", t.toString());
            }
        });
    }
}