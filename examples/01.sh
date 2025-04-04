// ======================================================== //
//         !!!Custom language is up and running!!!          //
// ======================================================== //

let input_file_name = "./test.txt"

// iterate over values and populate a hash map //
let iter_arr = df(arr, temp_arr, hmap, hmap_id, idx) {
    if (idx == len(arr)) { hmap }
    else {
        print(arr[idx]);
        if (arr[idx][0] != NULL) {
            let temp_arr = push(temp_arr, arr[idx][0]);
        }
        else {
            hmap[hmap_id] = temp_arr;
            let hmap_id = hmap_id + 1;
            let temp_arr = [];

        }
        iter_arr(input, temp_arr, hmap, hmap_id, idx + 1);
    }
}


// read in a files contents into hash map//
let input = read_file(input_file_name, " ", "INT");
let temp_arr = [];
let hmap = {}; 
let hmap_id = 0; 
let idx = 0;
let hmap = iter_arr(input, temp_arr, hmap, hmap_id, idx);

// ======================================================== //
// now sum all sub arrays to find answer AoC 2022 day 1!!!  //
// ======================================================== //

// define a sum function with builtin array functions//
let reduce = df(arr, initial, f) {
    let iter = df(arr, result) {
        if (len(arr) == 0) {
            result
        } else {
            iter(rest(arr), f(result, first(arr)));
        }
    };
    iter(arr, initial);
};

let sum = df(arr) {
    reduce(arr, 0, df(initial, el) { initial + el });
};

// iterate over the entire hash and find the key with the largest sum //
let iter_hash = df(hash, max_sum, key) {
    if (key == len(hash)) { max_sum }
    else {
        let x = sum(hash[key]);
        if (x > max_sum){let max_sum = x;}
        iter_hash(hash, max_sum, key + 1);
    }
}

let max_sum = 0;
let max_sum = iter_hash(hmap, max_sum, 0);
print("Awesome we solved AoC day 1 from 2022!")
print(max_sum);
