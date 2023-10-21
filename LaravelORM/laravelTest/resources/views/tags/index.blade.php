@extends('layouts.app')
@section('content')
<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-8">
            

            <div class="card">
                @foreach ($tags as $tag)
                <h2>{{$tag->title}}</h2>
                <ul> 
                    @foreach ($tag->posts as $post)
                     <li>{{$post->title}}</li>
                    @endforeach
                </ul>

                <h2>****</h2>
                @endforeach
            </div>
        </div>
    </div>
</div>
@endsection