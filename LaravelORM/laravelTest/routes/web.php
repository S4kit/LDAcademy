<?php

use Illuminate\Support\Facades\Route;
use App\Models\Adress;
use App\Models\Post;
use App\Models\User;
use App\Models\Tag;




/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "web" middleware group. Make something great!
|
*/

Route::get('/', function () {
    return view('welcome');
});
Auth::routes();

Route::get('/home', [App\Http\Controllers\HomeController::class, 'index'])->name('home');
Route::get('/user', function () {
    $users = User::get();
    return view('users.index', compact('users'));

});
Route::get('/adress', function () {
    $adresses = Adress::get();
    return view('adresses.index', compact('adresses'));
});

Route::get('/posts', function () {


    //$posts = Post::with(['user', 'tags'])->get();
    //return view('posts.index', compact('posts'));
    //$tag = Tag::first();
    //$post = Post::first();
    //$post->tags()->attach([2, 3, 4]);

});

Route::get('/tags', function () {
    $tags = Tag::with('posts')->get();
    return view('tags.index', compact('tags'));

});