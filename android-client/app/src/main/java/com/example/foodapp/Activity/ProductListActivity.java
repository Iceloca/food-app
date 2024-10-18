package com.example.foodapp.Activity;

import android.graphics.Color;
import android.os.Bundle;
import android.view.View;
import android.view.Window;
import android.view.WindowManager;

import androidx.recyclerview.widget.GridLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.example.foodapp.Adapter.ProductListAdapter;
import com.example.foodapp.Domain.Product.ProductResponse;
import com.example.foodapp.databinding.ActivityProductListBinding;

import java.util.ArrayList;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

public class ProductListActivity extends BaseActivity {

    ActivityProductListBinding binding;
    private Integer categoryId;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        binding = ActivityProductListBinding.inflate(getLayoutInflater());

        Window window = getWindow();
        window.clearFlags(WindowManager.LayoutParams.FLAG_TRANSLUCENT_STATUS);
        window.addFlags(WindowManager.LayoutParams.FLAG_DRAWS_SYSTEM_BAR_BACKGROUNDS);
        window.setStatusBarColor(Color.WHITE);
        View decorView = window.getDecorView();
        decorView.setSystemUiVisibility(View.SYSTEM_UI_FLAG_LIGHT_STATUS_BAR);
        setContentView(binding.getRoot());

        getIntentExtra();
        initListFood();
    }


    private void getIntentExtra() {
        categoryId = getIntent().getIntExtra("CategoryID", 0);
        String categoryName = getIntent().getStringExtra("CategoryName");
        binding.nameCategoryList.setText(categoryName);
        binding.returnButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                finish();
            }
        });

    }

    private void initListFood() {
        Call<ArrayList<ProductResponse>> call = productApi.GetProductsByCategory(categoryId);
        binding.progressListActivity.setVisibility(View.VISIBLE);
        call.enqueue(new Callback<ArrayList<ProductResponse>>() {
            @Override
            public void onResponse(Call<ArrayList<ProductResponse>> call, Response<ArrayList<ProductResponse>> response) {
                if (response.isSuccessful() && response.body() != null) {
                    binding.productListRecyclerView.setLayoutManager(new GridLayoutManager(ProductListActivity.this, 2));
                    RecyclerView.Adapter<ProductListAdapter.viewholder> adapter = new ProductListAdapter(response.body());
                    binding.productListRecyclerView.setAdapter(adapter);
                    binding.progressListActivity.setVisibility(View.GONE);
                }
            }

            @Override
            public void onFailure(Call<ArrayList<ProductResponse>> call, Throwable t) {

            }
        });
    }
}