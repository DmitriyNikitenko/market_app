package com.marketapp.user.service;

import com.marketapp.user.dto.UpdateProfileRequest;
import com.marketapp.user.dto.UserDto;
import com.marketapp.user.entity.UserProfile;
import com.marketapp.user.repository.UserProfileRepository;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class UserService {

    private final UserProfileRepository userProfileRepository;

    public UserService(UserProfileRepository userProfileRepository) {
        this.userProfileRepository = userProfileRepository;
    }

    @Transactional
    public UserDto getOrCreate(Long userId, String email) {
        UserProfile profile = userProfileRepository.findById(userId)
                .orElseGet(() -> userProfileRepository.save(new UserProfile(userId, email)));
        return toDto(profile);
    }

    @Transactional
    public UserDto updateProfile(Long userId, String email, UpdateProfileRequest request) {
        UserProfile profile = userProfileRepository.findById(userId)
                .orElseGet(() -> new UserProfile(userId, email));

        if (request.getFirstName() != null) {
            profile.setFirstName(request.getFirstName());
        }
        if (request.getLastName() != null) {
            profile.setLastName(request.getLastName());
        }
        if (request.getPhone() != null) {
            profile.setPhone(request.getPhone());
        }

        UserProfile saved = userProfileRepository.save(profile);
        return toDto(saved);
    }

    @Transactional(readOnly = true)
    public UserDto getById(Long id) {
        UserProfile profile = userProfileRepository.findById(id)
                .orElseThrow(() -> new IllegalArgumentException("Profile with ID " + id + " not found"));
        return toDto(profile);
    }

    private UserDto toDto(UserProfile profile) {
        return new UserDto(
                profile.getId(),
                profile.getEmail(),
                profile.getFirstName(),
                profile.getLastName(),
                profile.getPhone()
        );
    }
}