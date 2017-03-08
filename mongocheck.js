var found =false
db.locks.find({"acquired" : {$lt : new Date(new Date() - 1*1000)}}).forEach(function(n) {
// printjson(n); 
 found=true
});

if (found==true){
                quit(0);
}else{
                quit(2);
}