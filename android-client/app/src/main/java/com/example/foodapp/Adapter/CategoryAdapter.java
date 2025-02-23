package com.example.foodapp.Adapter;

import android.annotation.SuppressLint;
import android.content.Context;
import android.content.Intent;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.example.foodapp.Activity.ProductListActivity;
import com.example.foodapp.Domain.Product.CategoryResponse;
import com.example.foodapp.R;

import org.jetbrains.annotations.NotNull;

import java.util.ArrayList;

public class CategoryAdapter extends RecyclerView.Adapter<CategoryAdapter.viewholder>{
    private final ArrayList<CategoryResponse> items;
    private Context context;

    public CategoryAdapter(ArrayList<CategoryResponse> items) {
        this.items = items;
    }

    @NonNull
    @Override
    public CategoryAdapter.viewholder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        context = parent.getContext();
        View inflate = LayoutInflater
                .from(parent.getContext())
                .inflate(R.layout.viewholder_category, parent, false);
        return new viewholder(inflate);
    }

    @Override
    public void onBindViewHolder(@NonNull CategoryAdapter.viewholder holder, @SuppressLint("RecyclerView") int position) {
        holder.categoryName.setText(items.get(position).getName());

        Glide.with(context)
               .load(items.get(position).getImageURL())
               .into(holder.pic);

        holder.itemView.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                Intent intent = new Intent(context, ProductListActivity.class);
                intent.putExtra("CategoryID", items.get(position).getId());
                intent.putExtra("CategoryName", items.get(position).getName());
                context.startActivity(intent);
            }
        });
    }

    @Override
    public int getItemCount() {
        return items.size();
    }

    public static class viewholder extends RecyclerView.ViewHolder {
        TextView categoryName;
        ImageView pic;

        public viewholder(@NotNull View itemView) {
            super(itemView);
            categoryName = itemView.findViewById(R.id.categoryName);
            pic = itemView.findViewById(R.id.categoryPic);
        }
    }
}
