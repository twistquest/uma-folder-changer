# usage guide:
click on the exe found on the releases page  
it switches  
thats it :D

the above assumes that you already have the folder structure set up; this means that you have 2 folders named either "Umamusumegb" and "Umamusume" or "Umamusumejp" and "Umamusume". if they are not called exactly that the program will not do anything. the second prerequisite is therefore having both games downloaded so the folders mentioned above are correctly populated with the data needed. 

the reason for this program existing in the first place is that as of right now the global and japanese versions of umamusume occupy the same exact subfolder in the AppData/LocalLow/Cygames folder. this in practice means that everytime you want to run the other version (f.ex global), the game tries to access the folder which now contains data from the other version (in this situation JP) and errors out because it doesent see what it expects. a workaround for this is to rename one versions folder when youre launching the other, so that you can have both regions folders intact. this program is meant to simplify that process by doing the renaming for you, so you don't have to navigate to appdata each time
