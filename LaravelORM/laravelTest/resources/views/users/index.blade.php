@extends('layouts.app')
@section('content')
<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-8">
            

            <div class="card">
                @foreach ($users as $user)
                <h2>{{$user->name}}</h2>
                    @foreach ($user->adresses as $adress)
                    <h2>{{$adress->country}}</h2>
                        @foreach($user->posts as $post)
                        <h2>{{$post->title}}</h2>
                        @endforeach
                    @endforeach
                <h2>****</h2>
                @endforeach
            </div>
        </div>
    </div>
</div>
@endsection