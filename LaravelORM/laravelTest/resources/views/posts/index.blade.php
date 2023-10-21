@extends('layouts.app')
@section('content')
<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-8">
            

            <div class="card">
                @foreach ($posts as $post)
                <h2>{{$post->title}}</h2>
                <h2>{{$post->user->name}}</h2>
                <ul> 
                    @foreach ($post->tags as $tag)
                     <li>{{$tag->title}}</li>
                    @endforeach
                </ul>
                <h2>****</h2>
                @endforeach
            </div>
        </div>
    </div>
</div>
@endsection