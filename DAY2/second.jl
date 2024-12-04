function checkValidityTolerant(level)
    #qui devo solo controllare che si rispettino li standard dei livelli e, zappare via quelli sbagliati
    for i in 1:(size(level,1)-1)
        if i == 1
            if !(1 <= abs(level[2]-level[1]) <= 3)
                return checkValidity(deleteat!(level,i))
            end
        else
            lDiff = abs(level[i]-level[i-1])
            rDiff = abs(level[i]-level[i+1])
            if (!(1 <= lDiff <= 3) || !(1 <= rDiff <= 3) && ((level[i-1]<level[i]<level[i+1]) || ((level[i-1]>level[i]>level[i+1]))))
                return checkValidity(deleteat!(level,i))
            end
        end
    end
    return true
end


function checkValidity(lel)
    # first scan assets that the values are all
    # ASC
    asc = true
    desc = true 
    direction = false
    l = lel[1]
    for i in 2:(size(lel,1))
        if lel[i] < l 
            l = lel[i]
        else
            asc = false
            break
        end
    end
    #DESC
    l = lel[1]
    for i in 2:size(lel,1)
        if lel[i] > l
            l = lel[i]
        else
            desc = false
            break
        end
    end
    if (asc || desc) == false
        return false
    end

    for (i,l) in enumerate(lel)
        if i == 1
            if !(1 <= abs(lel[2]-lel[1]) <= 3)
                return false
            end
        elseif i == size(lel,1)
            if !(1 <= abs(lel[i-1]-lel[i]) <= 3)
                return false
            end
        else
            lDiff = abs(lel[i]-lel[i-1])
            rDiff = abs(lel[i]-lel[i+1])
            if !(1 <= lDiff <= 3) || !(1 <= rDiff <= 3)
                return false
            end
        end
    end
    return true
end
function main()
    resp = 0
    open("input.txt","r") do f
        for line in eachline(f)
            singleLevel = split(line," ")
            lel = Int64[]
            for el in singleLevel
                parsed = parse(Int,el)
                append!(lel,parsed)
            end
            if checkValidityTolerant(lel)
                resp = resp + 1
            end
        end
    end
    println(resp)
end

main()


