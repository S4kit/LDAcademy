@extends('layouts.app')
@section('content')
<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-8">
            

            <div class="card">
                @foreach ($adresses as $adress)
                <h2>{{$adress->id}}</h2>
                <h2>{{$adress->country}}</h2>
                <h2>{{$adress->user->id}}</h2>
                <h2>******</h2>
                @endforeach
            </div>
        </div>
    </div>
</div>
@endsection